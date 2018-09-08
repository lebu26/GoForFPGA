# Exploring Go as an FPGA programming language

This project is to investegate the performance of the novel platform [Reconfigure.io](https://reconfigure.io), which convert go to RTL language for [AWS f1 instance](https://aws.amazon.com/ec2/instance-types/f1/). Meanwhile, discover the optimization design patterns for this novel platform. The initial investegate indicates that compared with Xilinx SDAccel development environment(OpenCL programming interface), the Reco platform is inferior in default optimization performance. The latency of the program and the resource utilization are both higher than SDAccel. Therefore, 3 different design patterns are proposed to optimize these banchemark algorithms so that it can realize a twofold latency improvement compare with SDAccel. Corresponding tools are developed in C++ to help the process of this optimization.

This github Repo will includes all projects have done on Reco and SDAccel for examing and utilization. The testbench forlder contains testbench projects including Vector Product, FIR filter and Julia set. The tools folder contains the the tool for generating the FIR filter, automating loop unroll, and makeing Thread pool pattern optimzation.

The details of the optimization method can be found in the paper. 

## How to run the Reco project

Firstly, you will need to configure the Reco environment in your local device(A linux system is recommanded) following the [instruction](http://docs.reconfigure.io/overview.html) provided by Reco. Then, you can create new projects using the code in the corresponding project folder and do corrsponding deployment, simulation, and analysis. 

For professors from imperial college, you can use `reco auth <API key>` to log on my account to check these projects, where API key is the key I emailed to you.

For example, if you want to check the result of the data transformation:


1, Use `reco project set fir_data_time` to navigate to the project. After this step, you might want to use `reco project ls` to check whether this project is set(This step is optional). If the project is set, you can check what I have done in this project by check logs. Or you can also make your own deployment accroding to the [reco document](http://docs.reconfigure.io/overview.html).


2, For example, if you want to check the deployment result of the data tranformation benchmark. After the project is set, you can use `reco deploy ls` to check the deployment list. following message will come out. The deployment command test_fir 10000000 will include the experiment result of 10 million data transformation result
```
             DEPLOYMENT ID                                BUILD ID                         COMMAND             STATUS          STARTED        DURATION
  80f3eb58-3f65-4fae-909a-5735ab6b9462      6cc61c2d-d6ec-4a8f-ba0a-25bb12dda8b1      test_fir 10000000       completed      1 month ago      2m30s
  ac00b97b-41ab-45f9-b690-75bd40739ed2      6cc61c2d-d6ec-4a8f-ba0a-25bb12dda8b1      test_fir 100000000      completed      1 month ago      2m34s
 ```
 3, You can use `reco deploy log <deployment ID>` to check the deployment result, where deployment ID is the ID in the list. The result is like this deployment log of "80f3eb58-3f65-4fae-909a-5735ab6b9462".
 ```
 2018/08/05 12:12:38 **********************program starting*********************

2018/08/05 12:12:38 The arg is 10000000

2018/08/05 12:12:38 The number of test cases is 10000000

2018/08/05 12:12:39 hardware part staring...

Device/Slot[0] (/dev/xdma0, 0:0:1d.0)
xclProbe found 1 FPGA slots with XDMA driver running
2018/08/05 12:12:45 kernel arguments initialization finished...

2018/08/05 12:12:45 *************************************************************************************

2018/08/05 12:12:45                                   kernel staring

2018/08/05 12:12:45 *************************************************************************************


2018/08/05 12:12:45 **********************hardware finished*********************
2018/08/05 12:12:45 ** takes = 0.642861364 s

2018/08/05 12:12:45 *************************************************************************************

2018/08/05 12:12:45                                   kernel finished successfully

2018/08/05 12:12:45 *************************************************************************************
 ```
4, Other API you might use to expore this project can be found on this [ReconfigureIO page](http://docs.reconfigure.io/tool_invocation.html). 

For example, you might use this tool to explore 

* the teak diagram by 
```
reco graph ls
reco graph open <grpah ID>
```
* the resource utilization report by
```
reco build ls
reco build report <build ID>
```
* the build report that contains clock frquency message 
```
reco build ls
reco build log <build ID>
```
This table includes all the directory names, corrsponding projects name and the description of the project. You can find corresponding  source code, build result, teak graph in corresponding folder. These project names can be used to navigate to the project use `reco prject set <project name>`. 

<table><tr><th colspan="4">Directory name</th><th>Decirption</th><th>project name in my account</th></tr><tr><td rowspan="20">testbench</td><td rowspan="2">data tranformation</td><td>Reco</td><td></td><td>data tranformation&nbsp;&nbsp;&nbsp;on Reco platform</td><td>fir_data_time</td></tr><tr><td>SDAccel</td><td></td><td>data tranformation&nbsp;&nbsp;&nbsp;on SDAccel platform</td><td></td></tr><tr><td rowspan="7">VectorDp</td><td rowspan="6">Reco</td><td>n20-non-opti</td><td>vector product with 20 support vectors on Reco platform with no optimization applied</td><td>svm_n20sv</td></tr><tr><td>n1_Dpipe</td><td>vector product with 1 support vector on Reco platform with default-pipeline design pattern</td><td>svm_self_pipe_1</td></tr><tr><td>n15_Dpipe</td><td>vector product with 15 support vectors on Reco platform with default-pipeline design pattern</td><td>svm_2_another_pipe</td></tr><tr><td>n20_Dpipe</td><td>vector product with 20 support vectors on Reco platform with default-pipeline design pattern</td><td>svm_20sv_self_pipe</td></tr><tr><td>n20_Upipe</td><td>vector product with 20 support vector on Reco platform with user-defined-pipeline design pattern</td><td>svm_20_user_pipe</td></tr><tr><td>data time</td><td>data tranformation time for the vector product algorithm on Reco</td><td>svm_data</td></tr><tr><td>SDAccel</td><td>pipeline</td><td>vector product with 20 support vectors on SDAccel platform with default optimization applied</td><td></td></tr><tr><td rowspan="6">FIR</td><td rowspan="4">Reco</td><td>non-opti</td><td>FIR filter of order 10 on Reco with default optimization</td><td>fir_no_opti</td></tr><tr><td>n10_pipe</td><td>FIR filter of order 10 on Reco with user-defined design pattern optimization</td><td>FIR_n10_pipe</td></tr><tr><td>n20_pipe</td><td>FIR filter of order 20 on Reco with user-defined design pattern optimization</td><td>FIR_n20_pipe</td></tr><tr><td>n30_pipe</td><td>FIR filter of order 30 on Reco with user-defined design pattern optimization</td><td>FIR_n30_pipe</td></tr><tr><td rowspan="2">SDAccel</td><td>no opti</td><td>FIR filter of order 10 on SDAccel with default optimization</td><td></td></tr><tr><td>pipeline</td><td>FIR filter of order 10 on SDAccel with pipeline optimization</td><td></td></tr><tr><td rowspan="5">Julia Set</td><td rowspan="3">Reco</td><td>non-opti</td><td>julia set generator on Reco with default optimization</td><td>julia_32_no_opti</td></tr><tr><td>tp10</td><td>julia set generator on Reco with thread pool optimization with 10 threads</td><td>julia_32_pipe</td></tr><tr><td>tp20</td><td>julia set generator on Reco with thread pool optimization with 20 threads</td><td>julia_32_thread20</td></tr><tr><td rowspan="2">SDAccel</td><td>non-opti</td><td>julai set generator on SDAccel with default optimization</td><td></td></tr><tr><td>pipeline</td><td>julai set generator on SDAccel with pipeline optimization</td><td></td></tr><tr><td>tools</td><td>loopUnrollMaker</td><td></td><td></td><td>completely unroll the loop user provided with the loop body and iteration</td><td></td></tr><tr><td></td><td>fir_generator</td><td></td><td></td><td>generate functions for a&nbsp;&nbsp;pipelined FIR filter with reqired order, data type, and coefficients</td><td></td></tr><tr><td></td><td>TpMaker</td><td></td><td></td><td>generate functions for a thread pool with the number of thread and target function body.</td><td></td></tr></table>

## How to run the SDAccel project
Firstly You will need to follow [this project](https://github.com/Xilinx/SDAccel_Examples/wiki/Create,-configure-and-test-an-AWS-F1-instance) to configure the develop environment for SDAccel.

* Option 1: rebuild the project from the source code. 
After the Centos system is settled and the SDx GUI system is open. Create a new sdx project with the `xilinx_aws-vu9p-f1-04261818_dynamic_5_0.xpfm` platform and select the vector add example. After the vector add example project is created, copy the source code from this repository to replace the code in the example. Then build the project to get the xclbin file for FPGA and .exe file for cpu. Afterward, build the AFI use the xclbin file.

* Option2 : Build the AFI directly from the `xxx.xclbin` file. 
After the SDAccel enviroment is settled, upload the corresponding `xclbin` file and the `.exe` file to a folder. Then use the same intruction in this [page](https://github.com/Xilinx/SDAccel_Examples/wiki/Create,-configure-and-test-an-AWS-F1-instance) for building AFI and deployment.  

Finally, after the AFI is built, deploy to the FPGA according to the instruction in [page](https://github.com/Xilinx/SDAccel_Examples/wiki/Create,-configure-and-test-an-AWS-F1-instance) for deployment. Note that the "helloworld" command mentiond in the page should be replaced by the "xxx.exe" file.

## How to use tools for Reco optimization
The most significant contribution of this project are the design patterns proposed in the paper. These C++ tools are just used to duplicate the code in the case that mannually type the code is tedious.

* The DPipeMaker is used to make the default pipeline structure of vector product testbench with the fixed number of iteration.
* The loopUnrollMaker is used to completely unroll a loop with the fixed number of the iteration.
* The fir_generator will generate the FIR filter with the user-defined pipeline structure.
* The Tpmaker will generate a top level function for the thread poool pattern optimized program with the desired number of threads

The way of using these tool is to compile the `xxx.h` file with `xxx.cpp` file with gcc or IDE like visual studio. The corresponding arguments in the main function in the `xxx.cpp` file can be replaced by your program body and desired parameters.
