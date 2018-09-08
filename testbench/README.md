# Testbenches used for banchmark the ReconfigureIO

## Vector dot product
This alorithm is the dot product of a 16 element vetors with another set of vectors. 

The non-optimized version of this algorithm for 20 suppoert vector is conatined in [vectorDp/Reco/n20-non-opti](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/n20-non-opti) 

Meanwhile The same project on SDAccel is in [vectorDp/SDAccel/pipeline](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/SDAccel/pipeline)

It can be optimized by defualt pipeline design pattern, which eliminates the control flow by completely unroll the loop. The optimzed version using this pattern is included in folder name with "_Dpipe"

[With 1 support vector](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/n1_Dpipe)

[With 15 support vector](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/n15_Dpipe)

[With 20 support vector](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/n20_Dpipe)

This algorithm can also be optimized by the user defined pipeline design pattern, which utilizes the syntax-directed feather of the Reco platform to implement a pipline struceture. This structure is included in [vectorDp/Reco/n20_Upipe](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/n20_Upipe)

The data trandofrmation time of this algorithm is also investigated. The project is in [vectorDp/Reco/data time](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/data time)


## FIR
This is the Finite Impulse Response filter algorithm. The non optimized version of this algorithm is in [FIR/Reco/non-opti](https://github.com/lebu26/GoForFPGA/tree/master/testbench/FIR/Reco/non-opti)

Due to the strict data dependency in this algorithm, the default pipeline design pattern is not applicable. The optimization is realized through user defined pipeline. 

[order 10 filter](https://github.com/lebu26/GoForFPGA/tree/master/testbench/FIR/Reco/n10_pipe)

[order 20 filter](https://github.com/lebu26/GoForFPGA/tree/master/testbench/FIR/Reco/n20_pipe)

[order 30 filter](https://github.com/lebu26/GoForFPGA/tree/master/testbench/FIR/Reco/n30_pipe)

The source code of SDAccel project of order 10 FIR filter is also provided in [FIR/SDAccel](https://github.com/lebu26/GoForFPGA/tree/master/testbench/FIR/SDAccel)

## Julia Set
This folder contains experiments regrading julia set generator. 

[Non-optimizaed version](https://github.com/lebu26/GoForFPGA/tree/master/testbench/julia/Reco/no-opti)

[Thread pool with 10 threads](https://github.com/lebu26/GoForFPGA/tree/master/testbench/julia/Reco/tp10)

[Non-optimizaed version](https://github.com/lebu26/GoForFPGA/tree/master/testbench/julia/Reco/tp20)

The corresponding source code for SDAccel can be file in [/julia/SDAccel](https://github.com/lebu26/GoForFPGA/tree/master/testbench/julia/SDAccel)
