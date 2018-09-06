/*******************************************************************************
Vendor: Xilinx
Associated Filename: vadd.cpp
Purpose: SDAccel vector addition

*******************************************************************************
Copyright (C) 2017 XILINX, Inc.

This file contains confidential and proprietary information of Xilinx, Inc. and
is protected under U.S. and international copyright and other intellectual
property laws.

DISCLAIMER
This disclaimer is not a license and does not grant any rights to the materials
distributed herewith. Except as otherwise provided in a valid license issued to
you by Xilinx, and to the maximum extent permitted by applicable law:
(1) THESE MATERIALS ARE MADE AVAILABLE "AS IS" AND WITH ALL FAULTS, AND XILINX
HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS, EXPRESS, IMPLIED, OR STATUTORY,
INCLUDING BUT NOT LIMITED TO WARRANTIES OF MERCHANTABILITY, NON-INFRINGEMENT, OR
FITNESS FOR ANY PARTICULAR PURPOSE; and (2) Xilinx shall not be liable (whether
in contract or tort, including negligence, or under any other theory of
liability) for any loss or damage of any kind or nature related to, arising under
or in connection with these materials, including for any direct, or any indirect,
special, incidental, or consequential loss or damage (including loss of data,
profits, goodwill, or any type of loss or damage suffered as a result of any
action brought by a third party) even if such damage or loss was reasonably
foreseeable or Xilinx had been advised of the possibility of the same.

CRITICAL APPLICATIONS
Xilinx products are not designed or intended to be fail-safe, or for use in any
application requiring fail-safe performance, such as life-support or safety
devices or systems, Class III medical devices, nuclear facilities, applications
related to the deployment of airbags, or any other applications that could lead
to death, personal injury, or severe property or environmental damage
(individually and collectively, "Critical Applications"). Customer assumes the
sole risk and liability of any use of Xilinx products in Critical Applications,
subject only to applicable laws and regulations governing limitations on product
liability.

THIS COPYRIGHT NOTICE AND DISCLAIMER MUST BE RETAINED AS PART OF THIS FILE AT
ALL TIMES.

*******************************************************************************/
#include <stdlib.h>
#include <fstream>
#include <iostream>
#include <time.h>
#include "vadd.h"

//TARGET_DEVICE macro needs to be passed from gcc command line
#if defined(SDX_PLATFORM) && !defined(TARGET_DEVICE)
    #define STR_VALUE(arg)      #arg
    #define GET_STRING(name) STR_VALUE(name)
    #define TARGET_DEVICE GET_STRING(SDX_PLATFORM)
#endif

static const int width = 100;
static const int height = 100;

static const std::string error_message =
    "Error: Result mismatch:\n"
    "i = %d CPU result = %d Device result = %d\n";

int mul(int a, int b){
    long la = (long)a;
    long lb = (long)b;
    long res = la*lb;
    return (int)((res+(1<<15))>>16);
}

int main(int argc, char* argv[]) {

    //TARGET_DEVICE macro needs to be passed from gcc command line
    const char *target_device_name = TARGET_DEVICE;

    if(argc != 2) {
        std::cout << "Usage: " << argv[0] <<" <xclbin>" << std::endl;
        return EXIT_FAILURE;
    }

    char* xclbinFilename = argv[1];
    
    // Compute the size of array in bytes
    size_t size_in_bytes = width*height * sizeof(int);
    
    // Creates a vector of DATA_SIZE elements with an initial value of 10 and 32
    // using customized allocator for getting buffer alignment to 4k boundary
    std::vector<int,aligned_allocator<int>> source_results(width*height);
    
    std::vector<cl::Device> devices;
    cl::Device device;
    std::vector<cl::Platform> platforms;
    bool found_device = false;

    //traversing all Platforms To find Xilinx Platform and targeted
    //Device in Xilinx Platform
    cl::Platform::get(&platforms);
    for(size_t i = 0; (i < platforms.size() ) & (found_device == false) ;i++){
        cl::Platform platform = platforms[i];
        std::string platformName = platform.getInfo<CL_PLATFORM_NAME>();
        if ( platformName == "Xilinx"){
            devices.clear();
            platform.getDevices(CL_DEVICE_TYPE_ACCELERATOR, &devices);

            //Traversing All Devices of Xilinx Platform
            for (size_t j = 0 ; j < devices.size() ; j++){
                device = devices[j];
                std::string deviceName = device.getInfo<CL_DEVICE_NAME>();
                if (deviceName == target_device_name){
                    found_device = true;
                    break;
                }
            }
        }
    }
    if (found_device == false){
       std::cout << "Error: Unable to find Target Device " 
           << target_device_name << std::endl;
       return EXIT_FAILURE; 
    }

    // Creating Context and Command Queue for selected device
    cl::Context context(device);
    cl::CommandQueue q(context, device, CL_QUEUE_PROFILING_ENABLE);

    // Load xclbin 
    std::cout << "Loading: '" << xclbinFilename << "'\n";
    std::ifstream bin_file(xclbinFilename, std::ifstream::binary);
    bin_file.seekg (0, bin_file.end);
    unsigned nb = bin_file.tellg();
    bin_file.seekg (0, bin_file.beg);
    char *buf = new char [nb];
    bin_file.read(buf, nb);
    
    // Creating Program from Binary File
    cl::Program::Binaries bins;
    bins.push_back({buf,nb});
    devices.resize(1);
    cl::Program program(context, devices, bins);
    
    // This call will get the kernel object from program. A kernel is an 
    // OpenCL function that is executed on the FPGA. 
    cl::Kernel krnl_vector_add(program,"krnl_vadd");
    
    // These commands will allocate memory on the Device. The cl::Buffer objects can
    // be used to reference the memory locations on the device. 
    cl::Buffer buffer_result(context, CL_MEM_USE_HOST_PTR | CL_MEM_WRITE_ONLY, 
            size_in_bytes, source_results.data());
    
    // Data will be transferred from system memory over PCIe to the FPGA on-board
    // DDR memory.
    //set the kernel Arguments
    int para_w = (int)((2.0/float(width))*(1<<16));
    int para_h = (int)((2.0/float(height))*(1<<16));
    int cX = (int)(-0.7*(1<<16));
    int cY = (int)(0.27*(1<<16));

    int narg=0;
    krnl_vector_add.setArg(narg++,buffer_result);
    krnl_vector_add.setArg(narg++,width);
    krnl_vector_add.setArg(narg++,height);
    krnl_vector_add.setArg(narg++,para_w);
    krnl_vector_add.setArg(narg++,para_h);
    krnl_vector_add.setArg(narg++,cX);
    krnl_vector_add.setArg(narg++,cY);

    //Launch the Kernel
    struct timespec tt1, tt2;
    clock_gettime(CLOCK_REALTIME,&tt1);
    q.enqueueTask(krnl_vector_add);

    // The result of the previous kernel execution will need to be retrieved in
    // order to view the results. This call will transfer the data from FPGA to
    // source_results vector
    q.enqueueMigrateMemObjects({buffer_result},CL_MIGRATE_MEM_OBJECT_HOST);
    q.finish();
    clock_gettime(CLOCK_REALTIME,&tt2);
    double diff_hw=double(tt2.tv_nsec - tt1.tv_nsec)/1000000000;
    double diff_sec_hw=double(tt2.tv_sec - tt1.tv_sec);

    //Verify the result
    struct timespec tt1_sw, tt2_sw;
    clock_gettime(CLOCK_REALTIME,&tt1_sw);
    int match = 0;
    int gold[width*height]={0};
    int f_1 = 65536; //1
    int f_1_5 = 98304; //1.5
    int f_2 = 131072;//2
    int f_4 = 262144;//4
    int maxIter = 255;
    for (int x = 0; x < width; x++) {
        int thisx = (int)(x*f_1);
        int tmp, zx, zy;
        for( int y = 0; y < height; y++) {
            int thisy = (int)(y*f_1);
            int tm = mul(thisx, para_w);
            tm = tm - f_1;
            zx = mul(f_1_5,tm);//f_1_5.Mul(thisx.Mul(para_w) - f_1)
            zy = mul(thisy,para_h) - f_1;
            int i = maxIter;

            int zx2 = mul(zx,zx);
            int zy2 = mul(zy,zy);

            while ((zx2+zy2) < f_4 && i > 0){
                //tmp = zx.Mul(zx) - zy.Mul(zy) - fixed.Int52_12(2867)
                tmp = zx2 - zy2 + cX;
                int tm1 = mul(zx,zy);
                zy = mul(f_2,tm1) + cY;
                zx = tmp;
                i--;
                zx2 = mul(zx,zx);
                zy2 = mul(zy,zy);
            }
            int index = x*height+y;
            gold[index]= i;
        }
    }

    clock_gettime(CLOCK_REALTIME,&tt2_sw);
    double diff_sw=(double)(tt2_sw.tv_nsec - tt1_sw.tv_nsec)/1000000000;
    double diff_sec_sw=(double)(tt2_sw.tv_sec - tt1_sw.tv_sec);
    int err_count = 0;
    for (int i = 0; i < height*width; i++) {
        int host_result = gold[i];
        if (source_results[i] != host_result) {
            printf(error_message.c_str(), i, host_result, source_results[i]);
            match = 1;
            break;
        }
        if(i<10){
            std::cout << i <<", "<<host_result<<std::endl;
        }
    }

    std::cout << "TEST " << (match ? "FAILED" : "PASSED") << std::endl;
     std::cout << "error count = " << err_count << std::endl;
     std::cout <<"hardware take " << diff_hw<<" ms and " << tt2.tv_sec - tt1.tv_sec<<"s, totally "<<diff_hw+diff_sec_hw<< "s. "<< std::endl;
    //printf("printf hardware takes %ld ns \n\r", tt2.tv_nsec-tt1.tv_sec);
     std::cout <<"software take " << diff_sw <<" ms and " << tt2_sw.tv_sec - tt1_sw.tv_sec<<"s, totally "<<diff_sw+diff_sec_sw<< "s. "<<std::endl;
    //printf("printf hardware takes %ld ns \n\r", tt2sw.tv_nsec-tt1sw.tv_sec);
     return (match ? EXIT_FAILURE :  EXIT_SUCCESS);

}
