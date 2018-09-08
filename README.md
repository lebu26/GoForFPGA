# Exploring Go as an FPGA programming language

This project is to investegate the performance of the novel platform [Reconfigure.io](https://reconfigure.io), which convert go to RTL language for [AWS f1 instance](https://aws.amazon.com/ec2/instance-types/f1/). Meanwhile, discover the optimization design patterns for this novel platform. The initial investegate indicates that compared with Xilinx SDAccel development environment(OpenCL programming interface), the Reco platform is inferior in default optimization performance. The latency of the program and the resource utilization are both higher than SDAccel. Therefore, 3 different design patterns are proposed to optimize these banchemark algorithms so that it can realize a twofold latency improvement compare with SDAccel. Corresponding tools are developed in C++ to help the process of this optimization.

This github Repo will includes all projects done on Reco and SDAccel for examing and utilization. The testbench forlder will contains testbench projects including Vector Product, FIR filter and Julia set. The tools folder will contains the the tool for generating the FIR filter, automating loop unroll, and makeing Thread pool pattern optimzation.

The details of the optimization method can be found in the paper. 

Reco platform uses [Teak](https://ieeexplore.ieee.org/document/5291078/) to perform the network level optimization. This optimization will result in a token-flow system with elasticity. Ideally, this network can work as a default pipeline that waiting for data with token coming, start the compute, and sent the token to next node. But this default optimization will be limited by the control flow of the program. Only when the control flow like for loops do not exist in the function, the pipeline structure is fully functional. Otherwise, the synthesised structure will stick to the logic described in the go code and the dataflow will be stucking at the the loop, waiting for the finish of last iteration.   

## Getting Started

To run the projects on Reco, you will need to configure the Reco environment in your local device following the [instruction](http://docs.reconfigure.io/overview.html) provided by Reco. 

Then, for my marker from imperial college, you can use 
```
reco auth <API key>
```
to log on my account to check these projects. Following table includes all the project names and corrsponding projects
<dl>
  <style type="text/css">
.tg  {border-collapse:collapse;border-spacing:0;}
.tg td{font-family:Arial, sans-serif;font-size:14px;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:black;}
.tg th{font-family:Arial, sans-serif;font-size:14px;font-weight:normal;padding:10px 5px;border-style:solid;border-width:1px;overflow:hidden;word-break:normal;border-color:black;}
.tg .tg-0pky{border-color:inherit;text-align:left;vertical-align:top}
</style>
<table class="tg">
  <tr>
    <th class="tg-0pky">Directory name</th>
    <th class="tg-0pky"></th>
    <th class="tg-0pky"></th>
    <th class="tg-0pky"></th>
    <th class="tg-0pky">Decirption</th>
    <th class="tg-0pky">project name in my account</th>
  </tr>
  <tr>
    <td class="tg-0pky">testbench</td>
    <td class="tg-0pky">data tranformation</td>
    <td class="tg-0pky">Reco</td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">data tranformation&nbsp;&nbsp;&nbsp;on Reco platform</td>
    <td class="tg-0pky">fir_data_time</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">SDAccel</td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">data tranformation&nbsp;&nbsp;&nbsp;on SDAccel platform</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">VectorDp</td>
    <td class="tg-0pky">Reco</td>
    <td class="tg-0pky">n20-non-opti</td>
    <td class="tg-0pky">vector product with 20 support vectors on Reco platform with no optimization applied</td>
    <td class="tg-0pky">svm_n20sv</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">n1_Dpipe</td>
    <td class="tg-0pky">vector product with 1 support vector on Reco platform with default-pipeline design pattern</td>
    <td class="tg-0pky">svm_self_pipe_1</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">n15_Dpipe</td>
    <td class="tg-0pky">vector product with 15 support vectors on Reco platform with default-pipeline design pattern</td>
    <td class="tg-0pky">svm_2_another_pipe</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">n20_Dpipe</td>
    <td class="tg-0pky">vector product with 20 support vectors on Reco platform with default-pipeline design pattern</td>
    <td class="tg-0pky">svm_20sv_self_pipe</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">n20_Upipe</td>
    <td class="tg-0pky">vector product with 20 support vector on Reco platform with user-defined-pipeline design pattern</td>
    <td class="tg-0pky">svm_20_user_pipe</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">data time</td>
    <td class="tg-0pky">data tranformation time for the vector product algorithm on Reco</td>
    <td class="tg-0pky">svm_data</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">SDAccel</td>
    <td class="tg-0pky">pipeline</td>
    <td class="tg-0pky">vector product with 20 support vectors on SDAccel platform with default optimization applied</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">FIR</td>
    <td class="tg-0pky">Reco</td>
    <td class="tg-0pky">non-opti</td>
    <td class="tg-0pky">FIR filter of order 10 on Reco with default optimization</td>
    <td class="tg-0pky">fir_no_opti</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">n10_pipe</td>
    <td class="tg-0pky">FIR filter of order 10 on Reco with user-defined design pattern optimization</td>
    <td class="tg-0pky">FIR_n10_pipe</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">n20_pipe</td>
    <td class="tg-0pky">FIR filter of order 20 on Reco with user-defined design pattern optimization</td>
    <td class="tg-0pky">FIR_n20_pipe</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">n30_pipe</td>
    <td class="tg-0pky">FIR filter of order 30 on Reco with user-defined design pattern optimization</td>
    <td class="tg-0pky">FIR_n30_pipe</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">SDAccel</td>
    <td class="tg-0pky">no opti</td>
    <td class="tg-0pky">FIR filter of order 10 on SDAccel with default optimization</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">pipeline</td>
    <td class="tg-0pky">FIR filter of order 10 on SDAccel with pipeline optimization</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">Julia Set</td>
    <td class="tg-0pky">Reco</td>
    <td class="tg-0pky">non-opti</td>
    <td class="tg-0pky">julia set generator on Reco with default optimization</td>
    <td class="tg-0pky">julia_32_no_opti</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">tp10</td>
    <td class="tg-0pky">julia set generator on Reco with thread pool optimization with 10 threads</td>
    <td class="tg-0pky">julia_32_pipe</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">tp20</td>
    <td class="tg-0pky">julia set generator on Reco with thread pool optimization with 20 threads</td>
    <td class="tg-0pky">julia_32_thread20</td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">SDAccel</td>
    <td class="tg-0pky">non-opti</td>
    <td class="tg-0pky">julai set generator on SDAccel with default optimization</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">pipeline</td>
    <td class="tg-0pky">julai set generator on SDAccel with pipeline optimization</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky">tools</td>
    <td class="tg-0pky">loopUnrollMaker</td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">completely unroll the loop user provided with the loop body and iteration</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">fir_generator</td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">generate functions for a&nbsp;&nbsp;pipelined FIR filter with reqired order, data type, and coefficients</td>
    <td class="tg-0pky"></td>
  </tr>
  <tr>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">TpMaker</td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky"></td>
    <td class="tg-0pky">generate functions for a thread pool with the number of thread and target function body.</td>
    <td class="tg-0pky"></td>
  </tr>
</table>
 <\dl>





### Prerequisites

What things you need to install the software and how to install them

```
Give examples
```

### Installing

A step by step series of examples that tell you how to get a development env running

Say what the step will be

```
Give the example
```

And repeat

```
until finished
```

End with an example of getting some data out of the system or using it for a little demo

## Running the tests

Explain how to run the automated tests for this system

### Break down into end to end tests

Explain what these tests test and why

```
Give an example
```

### And coding style tests

Explain what these tests test and why

```
Give an example
```

## Deployment

Add additional notes about how to deploy this on a live system

## Built With

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - The web framework used
* [Maven](https://maven.apache.org/) - Dependency Management
* [ROME](https://rometools.github.io/rome/) - Used to generate RSS Feeds

## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags). 

## Authors

* **Billie Thompson** - *Initial work* - [PurpleBooth](https://github.com/PurpleBooth)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc
