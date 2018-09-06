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
#include "vadd.h"
#include <time.h>
//#include <chrono>

//TARGET_DEVICE macro needs to be passed from gcc command line
#if defined(SDX_PLATFORM) && !defined(TARGET_DEVICE)
    #define STR_VALUE(arg)      #arg
    #define GET_STRING(name) STR_VALUE(name)
    #define TARGET_DEVICE GET_STRING(SDX_PLATFORM)
#endif

static const int DATA_SIZE = 1000000;

static const std::string error_message =
    "Error: Result mismatch:\n"
    "i = %d CPU result = %d Device result = %d\n";

int main(int argc, char* argv[]) {

    //TARGET_DEVICE macro needs to be passed from gcc command line
    const char *target_device_name = TARGET_DEVICE;

    if(argc != 2) {
        std::cout << "Usage: " << argv[0] <<" <xclbin>" << std::endl;
        return EXIT_FAILURE;
    }

    char* xclbinFilename = argv[1];
    
    // Compute the size of array in bytes
    size_t size_in_bytes = DATA_SIZE * sizeof(int);
    
    // Creates a vector of DATA_SIZE elements with an initial value of 10 and 32
    // using customized allocator for getting buffer alignment to 4k boundary
    std::vector<int,aligned_allocator<int>> source_a(DATA_SIZE*16, 2);
    //std::vector<int,aligned_allocator<int>> source_b(DATA_SIZE, 32);
    std::vector<int,aligned_allocator<int>> source_results(DATA_SIZE);
    
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
    cl::Buffer buffer_a(context, CL_MEM_USE_HOST_PTR | CL_MEM_READ_ONLY,  
            size_in_bytes*16, source_a.data());
    //cl::Buffer buffer_b(context, CL_MEM_USE_HOST_PTR | CL_MEM_READ_ONLY,
    //        size_in_bytes, source_b.data());
    cl::Buffer buffer_result(context, CL_MEM_USE_HOST_PTR | CL_MEM_WRITE_ONLY, 
            size_in_bytes, source_results.data());
    
    // Data will be transferred from system memory over PCIe to the FPGA on-board
    // DDR memory.
    q.enqueueMigrateMemObjects({buffer_a},0/* 0 means from host*/);

    //set the kernel Arguments
    int narg=0;
    krnl_vector_add.setArg(narg++,buffer_a);
    //krnl_vector_add.setArg(narg++,buffer_b);
    krnl_vector_add.setArg(narg++,buffer_result);
    krnl_vector_add.setArg(narg++,DATA_SIZE);


    //auto start_hw = std::chrono::high_precission_clock::now();
    std::cout << "hardware start" << std::endl;
    //Launch the Kernel
	struct timespec tt1, tt2;
	clock_gettime(CLOCK_REALTIME,&tt1);

    q.enqueueTask(krnl_vector_add);
    //start_hw = clock() - start_hw;
    //double diff_hw = float(start_hw)/CLOCKS_PER_SEC;

    // The result of the previous kernel execution will need to be retrieved in
    // order to view the results. This call will write the data from the
    // buffer_result cl_mem object to the source_results vector
    q.enqueueMigrateMemObjects({buffer_result},CL_MIGRATE_MEM_OBJECT_HOST);
    q.finish();
	clock_gettime(CLOCK_REALTIME,&tt2);
	double diff_hw=double(tt2.tv_nsec - tt1.tv_nsec)/1000000000;
	double diff_sec_hw=double(tt2.tv_sec - tt1.tv_sec);
    //int temp = 0

    //Verify the result
    int match = 0;
    int err_count = 0;
    //auto start_sw = std::chrono::high_precission_clock::now();
    struct timespec tt1sw, tt2sw;
//start_sw = clock();
	clock_gettime(CLOCK_REALTIME,&tt1sw);
	int SV[320]={0, -97, -3, -38, 14, 34, -13, 32, 22, -33, -11, -59, -28, -14, 82, 86, -34, -97, -67, -67, -44, 34, 13, 55, -4, -33, -35, -28, -28, -14, 132, 7, -34, -58, -67, 17, -44, -123, 13, 32, 129, -33, -11, 217, -83, -14, 82, 7, -34, 38, -35, 45, -44, 3, -68, 32, 129, -84, -11, 63, -83, 27, 132, 7, -67, -58, -35, -67, -44, -28, 68, -85, 48, -7, 110, 2, -56, 110, -42, -31, -34, 18, -67, 17, -44, -123, -13, 32, 129, -33, -11, 186, -83, -14, 82, 7, -101, -136, -131, -123, -73, -91, -13, 8, 75, -33, -11, 156, -83, -14, 57, 86, -34, -20, -35, -38, -14, 3, 123, -15, 22, -33, 110, 2, -28, 152, -67, -31, 0, -97, -3, -38, 14, 3, -13, 8, 22, -33, -11, -59, -28, -14, 82, 86, 32, -39, -3, 74, 14, 3, 41, 126, 75, -33, -35, -28, -28, -14, 132, 86, 0, 0, -35, -10, -14, -28, -13, 102, 48, -33, -11, -28, -28, -14, 132, 86, -67, -97, -35, -38, -44, -28, 68, -85, 48, -33, 110, 32, -56, 110, -67, 7, 32, 18, 27, 17, -14, -28, -41, -15, -4, 147, 37, 125, -56, -97, -67, -31, 0, -97, -35, -38, -14, 34, -13, 32, 129, -58, -35, 32, -28, -14, 57, 7, 0, 18, -3, 17, -14, 3, -68, 32, 102, -84, -11, 63, -83, 27, 132, 7, -34, 0, -67, -10, -44, -123, 13, 32, 156, -33, -35, 186, -83, -14, 82, 7, 0, -78, -3, -95, -44, -60, 13, 8, 48, 44, 86, 32, -1, 27, -42, -71, 0, 18, -35, -10, -73, -186, 151, 8, -4, 95, 110, -28, -83, -14, -42, -71, 0, 38, -35, 17, -14, -28, -13, 102, 48, -33, -11, -28, -28, -14, 132, 86, -34, -58, -35, 17, -44, -123, 13, 32, 129, -33, -11, 186, -83, -14, 82, -31, };
	int Alpha[20] = {33, 33, -832, 33, 33, -832, -832, 33, 33, 33, 33, 33, 33, -166, 33, -665, 33, 33, 33, -99, };
	int res[DATA_SIZE]={0};
	for (int i=0; i<DATA_SIZE; i++){
			int sum = 0;
			int tm[16] = {0};
			for(int w=0; w<16;w++){
				tm[w] = source_a[i*16+w];
			}
			for (int j = 0; j < 20;j++){
				int dp = 0;
				for(int w=0; w<16;w++){
					dp+=tm[w]*SV[j*16+w];
				}
				dp = dp *2;
				sum += Alpha[j] *dp;
			}
			res[i]=sum;
	}
	clock_gettime(CLOCK_REALTIME,&tt2sw);
	for(int i=0;i<DATA_SIZE;i++){
        if (source_results[i] != res[i]) {
            if(err_count <10)
            {
                printf(error_message.c_str(), i, res[i], source_results[i]);
            }
            match = 1;
            err_count++;
            //break;
        }
	}

    //start_sw = clock() -start_sw;

    double diff_sw=double(tt2sw.tv_nsec - tt1sw.tv_nsec)/1000000000;
    double diff_sec_sw=double(tt2sw.tv_sec - tt1sw.tv_sec);

    //auto stop_sw = std::chrono::high_precission_clock::now();
    //auto duration_sw = std::chrono::duration_cast<microseconds>(stop - start);
    std::cout << "TEST " << (match ? "FAILED" : "PASSED") << std::endl; 
    std::cout << "error count = " << err_count << std::endl;
    std::cout <<"hardware take " << diff_hw<<" ms and " << tt2.tv_sec - tt1.tv_sec<<"s, totally "<<diff_hw+diff_sec_hw<< "s. "<< std::endl;
	//printf("printf hardware takes %ld ns \n\r", tt2.tv_nsec-tt1.tv_sec);
    std::cout <<"software take " << diff_sw <<" ms and " << tt2sw.tv_sec - tt1sw.tv_sec<<"s, totally "<<diff_sw+diff_sec_sw<< "s. "<<std::endl;
	//printf("printf hardware takes %ld ns \n\r", tt2sw.tv_nsec-tt1sw.tv_sec);
    return (match ? EXIT_FAILURE :  EXIT_SUCCESS);
}
