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

#define BUFFER_SIZE 256
kernel __attribute__((reqd_work_group_size(1, 1, 1)))
void krnl_vadd(
                global const int* a,
                global int* c,
                const int length)
{
    int SV[320]={0, -97, -3, -38, 14, 34, -13, 32, 22, -33, -11, -59, -28, -14, 82, 86, -34, -97, -67, -67, -44, 34, 13, 55, -4, -33, -35, -28, -28, -14, 132, 7, -34, -58, -67, 17, -44, -123, 13, 32, 129, -33, -11, 217, -83, -14, 82, 7, -34, 38, -35, 45, -44, 3, -68, 32, 129, -84, -11, 63, -83, 27, 132, 7, -67, -58, -35, -67, -44, -28, 68, -85, 48, -7, 110, 2, -56, 110, -42, -31, -34, 18, -67, 17, -44, -123, -13, 32, 129, -33, -11, 186, -83, -14, 82, 7, -101, -136, -131, -123, -73, -91, -13, 8, 75, -33, -11, 156, -83, -14, 57, 86, -34, -20, -35, -38, -14, 3, 123, -15, 22, -33, 110, 2, -28, 152, -67, -31, 0, -97, -3, -38, 14, 3, -13, 8, 22, -33, -11, -59, -28, -14, 82, 86, 32, -39, -3, 74, 14, 3, 41, 126, 75, -33, -35, -28, -28, -14, 132, 86, 0, 0, -35, -10, -14, -28, -13, 102, 48, -33, -11, -28, -28, -14, 132, 86, -67, -97, -35, -38, -44, -28, 68, -85, 48, -33, 110, 32, -56, 110, -67, 7, 32, 18, 27, 17, -14, -28, -41, -15, -4, 147, 37, 125, -56, -97, -67, -31, 0, -97, -35, -38, -14, 34, -13, 32, 129, -58, -35, 32, -28, -14, 57, 7, 0, 18, -3, 17, -14, 3, -68, 32, 102, -84, -11, 63, -83, 27, 132, 7, -34, 0, -67, -10, -44, -123, 13, 32, 156, -33, -35, 186, -83, -14, 82, 7, 0, -78, -3, -95, -44, -60, 13, 8, 48, 44, 86, 32, -1, 27, -42, -71, 0, 18, -35, -10, -73, -186, 151, 8, -4, 95, 110, -28, -83, -14, -42, -71, 0, 38, -35, 17, -14, -28, -13, 102, 48, -33, -11, -28, -28, -14, 132, 86, -34, -58, -35, 17, -44, -123, 13, 32, 129, -33, -11, 186, -83, -14, 82, -31, };
    int Alpha[20] = {33, 33, -832, 33, 33, -832, -832, 33, 33, 33, 33, 33, 33, -166, 33, -665, 33, 33, 33, -99, };
    for (int i=0; i<length; i++){
            int sum = 0;
            int tm[16] = {0};
            for(int w=0; w<16;w++){
                tm[w] = a[i*16+w];
            }
            for (int j = 0; j < 20;j++){
                int dp = 0;
                for(int w=0; w<16;w++){
                    dp+=tm[w]*SV[j*16+w];
                }
                dp = dp *2;
                sum += Alpha[j] *dp;
            }
            c[i]=sum;
    }
}
