# Testbenches used for banchmark the ReconfigureIO

##Vector dot product
This alorithm is the dot product of 16 element vetors. The non-optimized version of this algorithm is conatined in [Reco/n20-non-opti](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/n20-non-opti) 

It can be optimized by defualt pipeline design pattern, which eliminates the control flow by completely unroll the loop. The optimzed version using this pattern is included in folder name with "_Dpipe"

It can also be optimized by the user defined pipeline design pattern, which utilizes the syntax-directed feather of the Reco platform to implement
a pipline struceture. This structure is included in [Reco/n20_Upipe](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/n20_Upipe)

The data trandofrmation time of this algorithm is also investigated. The project is in [Reco/data time](https://github.com/lebu26/GoForFPGA/tree/master/testbench/vectorDp/Reco/data time)
##FIR
This is the Finite Impulse Response filter algorithm.

Due to the strict data dependency in this algorithm, the default pipeline design pattern is not applicable. 
The optimized
