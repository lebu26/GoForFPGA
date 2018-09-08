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

| project name | project description|
| --- | --- |
||The data tranfer time|




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