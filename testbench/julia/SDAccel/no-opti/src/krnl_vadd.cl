/*******************************************************************************
Vendor: Xilinx
Associated Filename: krnl_vadd.cl
Purpose: SDx vector addition example
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

//------------------------------------------------------------------------------
//
// kernel:  vadd
//
// Purpose: Demonstrate Vector Add in OpenCL
//
//#include "ap_fixed.h"
//#define BUFFER_SIZE 256
int mul(int a, int b){
    long la = (long)a;
    long lb = (long)b;
    long res = la*lb;
    return (int)((res+(1<<15))>>16);
}

kernel __attribute__((reqd_work_group_size(1, 1, 1)))
void krnl_vadd(
                global int* c,
                const int width,
                const int height,
                const int para_w,
                const int para_h,
                const int cX,
                const int cY)
{
    //ap_fixed<64,52> para_w = ap_fixed<64,52>(para_w_i);
    //ap_fixed<64,52> para_h = ap_fixed<64,52>(para_h_i);
    //ap_fixed<64,52> cX = ap_fixed<64,52>(cX_i);
    //ap_fixed<64,52> cY = ap_fixed<64,52>(cY_i);
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
            c[index]= i;
        }
    }
}
