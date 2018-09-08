# Exploring Go as an FPGA programming language

This project is to investegate the performance of the novel platform [Reconfigure.io](https://reconfigure.io), which convert go to RTL language for [AWS f1 instance](https://aws.amazon.com/ec2/instance-types/f1/). Meanwhile, discover the optimization design patterns for this novel platform. The initial investegate indicates that compared with Xilinx SDAccel development environment(OpenCL programming interface), the Reco platform is inferior in default optimization performance. The latency of the program and the resource utilization are both higher than SDAccel. Therefore, 3 different design patterns are proposed to optimize these banchemark algorithms so that it can realize a twofold latency improvement compare with SDAccel. Corresponding tools are developed in C++ to help the process of this optimization.

This github Repo will includes all projects done on Reco and SDAccel for examing and utilization. The testbench forlder will contains testbench projects including Vector Product, FIR filter and Julia set. The tools folder will contains the the tool for generating the FIR filter, automating loop unroll, and makeing Thread pool pattern optimzation.

The details of the optimization method can be found in the paper. 

Reco platform uses [Teak](https://ieeexplore.ieee.org/document/5291078/) to perform the network level optimization. This optimization will result in a token-flow system with elasticity. Ideally, this network can work as a default pipeline that waiting for data with token coming, start the compute, and sent the token to next node. But this default optimization will be limited by the control flow of the program. Only when the control flow like for loops do not exist in the function, the pipeline structure is fully functional. Otherwise, the synthesised structure will stick to the logic described in the go code and the dataflow will be stucking at the the loop, waiting for the finish of last iteration. 

## Getting Started

To run the projects on Reco, you will need to configure the Reco environment in your local device following the [instruction](http://docs.reconfigure.io/overview.html) provided by Reco. Then, you can create new projects using the code in the corresponding project folder and do corrsponding deployment, simulation, and analysis. 

For professors from imperial college, you can use `reco auth <API key>` to log on my account to check these projects, where API key is the key I emailed to you.

For example, if you want to check the result of the data transformation:
1, Use `reco project set fir_data_time` to navigate to the project. After this step, you might want to use `reco project ls` to check whether this project is set(This step is optional). If the project is set, you can check what I have done in this project by check logs. Or you can also make your own deployment accroding to the [reco document](http://docs.reconfigure.io/overview.html).
2, For example, if you want to check the deployment result of the data tranformation benchmark. After the project is set, you can use `reco deploy ls` to check the deployment list. follow message will come out. The deployment command test_fir 10000000 will include the experiment result of 10 million data transformation result
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
4, Other API you might use to expore this project can be found on this [ReconfigureIO page](http://docs.reconfigure.io/tool_invocation.html). You might use this tool to explor the teak diagram, deployment log, or even deploy some project. 

This table includes all the directory names, corrsponding projects name and the description of the project. You can find corresponding  source code, build result, teak graph in corresponding folder. These project names can be used to navigate to the project use `reco prject set <project name>`. 

<table><tr><th colspan="4">Directory name</th><th>Decirption</th><th>project name in my account</th></tr><tr><td rowspan="20">testbench</td><td rowspan="2">data tranformation</td><td>Reco</td><td></td><td>data tranformation&nbsp;&nbsp;&nbsp;on Reco platform</td><td>fir_data_time</td></tr><tr><td>SDAccel</td><td></td><td>data tranformation&nbsp;&nbsp;&nbsp;on SDAccel platform</td><td></td></tr><tr><td rowspan="7">VectorDp</td><td rowspan="6">Reco</td><td>n20-non-opti</td><td>vector product with 20 support vectors on Reco platform with no optimization applied</td><td>svm_n20sv</td></tr><tr><td>n1_Dpipe</td><td>vector product with 1 support vector on Reco platform with default-pipeline design pattern</td><td>svm_self_pipe_1</td></tr><tr><td>n15_Dpipe</td><td>vector product with 15 support vectors on Reco platform with default-pipeline design pattern</td><td>svm_2_another_pipe</td></tr><tr><td>n20_Dpipe</td><td>vector product with 20 support vectors on Reco platform with default-pipeline design pattern</td><td>svm_20sv_self_pipe</td></tr><tr><td>n20_Upipe</td><td>vector product with 20 support vector on Reco platform with user-defined-pipeline design pattern</td><td>svm_20_user_pipe</td></tr><tr><td>data time</td><td>data tranformation time for the vector product algorithm on Reco</td><td>svm_data</td></tr><tr><td>SDAccel</td><td>pipeline</td><td>vector product with 20 support vectors on SDAccel platform with default optimization applied</td><td></td></tr><tr><td rowspan="6">FIR</td><td rowspan="4">Reco</td><td>non-opti</td><td>FIR filter of order 10 on Reco with default optimization</td><td>fir_no_opti</td></tr><tr><td>n10_pipe</td><td>FIR filter of order 10 on Reco with user-defined design pattern optimization</td><td>FIR_n10_pipe</td></tr><tr><td>n20_pipe</td><td>FIR filter of order 20 on Reco with user-defined design pattern optimization</td><td>FIR_n20_pipe</td></tr><tr><td>n30_pipe</td><td>FIR filter of order 30 on Reco with user-defined design pattern optimization</td><td>FIR_n30_pipe</td></tr><tr><td rowspan="2">SDAccel</td><td>no opti</td><td>FIR filter of order 10 on SDAccel with default optimization</td><td></td></tr><tr><td>pipeline</td><td>FIR filter of order 10 on SDAccel with pipeline optimization</td><td></td></tr><tr><td rowspan="5">Julia Set</td><td rowspan="3">Reco</td><td>non-opti</td><td>julia set generator on Reco with default optimization</td><td>julia_32_no_opti</td></tr><tr><td>tp10</td><td>julia set generator on Reco with thread pool optimization with 10 threads</td><td>julia_32_pipe</td></tr><tr><td>tp20</td><td>julia set generator on Reco with thread pool optimization with 20 threads</td><td>julia_32_thread20</td></tr><tr><td rowspan="2">SDAccel</td><td>non-opti</td><td>julai set generator on SDAccel with default optimization</td><td></td></tr><tr><td>pipeline</td><td>julai set generator on SDAccel with pipeline optimization</td><td></td></tr><tr><td>tools</td><td>loopUnrollMaker</td><td></td><td></td><td>completely unroll the loop user provided with the loop body and iteration</td><td></td></tr><tr><td></td><td>fir_generator</td><td></td><td></td><td>generate functions for a&nbsp;&nbsp;pipelined FIR filter with reqired order, data type, and coefficients</td><td></td></tr><tr><td></td><td>TpMaker</td><td></td><td></td><td>generate functions for a thread pool with the number of thread and target function body.</td><td></td></tr></table>

