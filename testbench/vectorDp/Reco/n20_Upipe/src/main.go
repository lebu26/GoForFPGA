package main

import (
	// Import the entire framework for interracting with SDAccel from Go (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	
  // Use the new AXI protocol package for interracting with memory
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"

	"github.com/ReconfigureIO/fixed"
)
type data struct {
	t0 fixed.Int26_6
	t1 fixed.Int26_6
	t2 fixed.Int26_6
	t3 fixed.Int26_6
	t4 fixed.Int26_6
	t5 fixed.Int26_6
	t6 fixed.Int26_6
	t7 fixed.Int26_6 
	t8 fixed.Int26_6
	t9 fixed.Int26_6
	t10 fixed.Int26_6
	t11 fixed.Int26_6 
	t12 fixed.Int26_6
	t13 fixed.Int26_6
	t14 fixed.Int26_6
	t15 fixed.Int26_6  
}

func svm0(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * -97
		dp += tm.t2 * -3
		dp += tm.t3 * -38
		dp += tm.t4 * 14
		dp += tm.t5 * 34
		dp += tm.t6 * -13
		dp += tm.t7 * 32
		dp += tm.t8 * 22
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * -59
		dp += tm.t12 * -28
		dp += tm.t13 * -14
		dp += tm.t14 * 82
		dp += tm.t15 * 86
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm1(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -34
		dp += tm.t1 * -97
		dp += tm.t2 * -67
		dp += tm.t3 * -67
		dp += tm.t4 * -44
		dp += tm.t5 * 34
		dp += tm.t6 * 13
		dp += tm.t7 * 55
		dp += tm.t8 * -4
		dp += tm.t9 * -33
		dp += tm.t10 * -35
		dp += tm.t11 * -28
		dp += tm.t12 * -28
		dp += tm.t13 * -14
		dp += tm.t14 * 132
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm2(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -34
		dp += tm.t1 * -58
		dp += tm.t2 * -67
		dp += tm.t3 * 17
		dp += tm.t4 * -44
		dp += tm.t5 * -123
		dp += tm.t6 * 13
		dp += tm.t7 * 32
		dp += tm.t8 * 129
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * 217
		dp += tm.t12 * -83
		dp += tm.t13 * -14
		dp += tm.t14 * 82
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(-832 * dp)
	}
}
func svm3(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -34
		dp += tm.t1 * 38
		dp += tm.t2 * -35
		dp += tm.t3 * 45
		dp += tm.t4 * -44
		dp += tm.t5 * 3
		dp += tm.t6 * -68
		dp += tm.t7 * 32
		dp += tm.t8 * 129
		dp += tm.t9 * -84
		dp += tm.t10 * -11
		dp += tm.t11 * 63
		dp += tm.t12 * -83
		dp += tm.t13 * 27
		dp += tm.t14 * 132
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm4(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -67
		dp += tm.t1 * -58
		dp += tm.t2 * -35
		dp += tm.t3 * -67
		dp += tm.t4 * -44
		dp += tm.t5 * -28
		dp += tm.t6 * 68
		dp += tm.t7 * -85
		dp += tm.t8 * 48
		dp += tm.t9 * -7
		dp += tm.t10 * 110
		dp += tm.t11 * 2
		dp += tm.t12 * -56
		dp += tm.t13 * 110
		dp += tm.t14 * -42
		dp += tm.t15 * -31
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm5(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -34
		dp += tm.t1 * 18
		dp += tm.t2 * -67
		dp += tm.t3 * 17
		dp += tm.t4 * -44
		dp += tm.t5 * -123
		dp += tm.t6 * -13
		dp += tm.t7 * 32
		dp += tm.t8 * 129
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * 186
		dp += tm.t12 * -83
		dp += tm.t13 * -14
		dp += tm.t14 * 82
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(-832 * dp)
	}
}
func svm6(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -101
		dp += tm.t1 * -136
		dp += tm.t2 * -131
		dp += tm.t3 * -123
		dp += tm.t4 * -73
		dp += tm.t5 * -91
		dp += tm.t6 * -13
		dp += tm.t7 * 8
		dp += tm.t8 * 75
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * 156
		dp += tm.t12 * -83
		dp += tm.t13 * -14
		dp += tm.t14 * 57
		dp += tm.t15 * 86
		dp = dp*2
		res <- uint32(-832 * dp)
	}
}
func svm7(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -34
		dp += tm.t1 * -20
		dp += tm.t2 * -35
		dp += tm.t3 * -38
		dp += tm.t4 * -14
		dp += tm.t5 * 3
		dp += tm.t6 * 123
		dp += tm.t7 * -15
		dp += tm.t8 * 22
		dp += tm.t9 * -33
		dp += tm.t10 * 110
		dp += tm.t11 * 2
		dp += tm.t12 * -28
		dp += tm.t13 * 152
		dp += tm.t14 * -67
		dp += tm.t15 * -31
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm8(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * -97
		dp += tm.t2 * -3
		dp += tm.t3 * -38
		dp += tm.t4 * 14
		dp += tm.t5 * 3
		dp += tm.t6 * -13
		dp += tm.t7 * 8
		dp += tm.t8 * 22
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * -59
		dp += tm.t12 * -28
		dp += tm.t13 * -14
		dp += tm.t14 * 82
		dp += tm.t15 * 86
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm9(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 32
		dp += tm.t1 * -39
		dp += tm.t2 * -3
		dp += tm.t3 * 74
		dp += tm.t4 * 14
		dp += tm.t5 * 3
		dp += tm.t6 * 41
		dp += tm.t7 * 126
		dp += tm.t8 * 75
		dp += tm.t9 * -33
		dp += tm.t10 * -35
		dp += tm.t11 * -28
		dp += tm.t12 * -28
		dp += tm.t13 * -14
		dp += tm.t14 * 132
		dp += tm.t15 * 86
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm10(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * 0
		dp += tm.t2 * -35
		dp += tm.t3 * -10
		dp += tm.t4 * -14
		dp += tm.t5 * -28
		dp += tm.t6 * -13
		dp += tm.t7 * 102
		dp += tm.t8 * 48
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * -28
		dp += tm.t12 * -28
		dp += tm.t13 * -14
		dp += tm.t14 * 132
		dp += tm.t15 * 86
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm11(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -67
		dp += tm.t1 * -97
		dp += tm.t2 * -35
		dp += tm.t3 * -38
		dp += tm.t4 * -44
		dp += tm.t5 * -28
		dp += tm.t6 * 68
		dp += tm.t7 * -85
		dp += tm.t8 * 48
		dp += tm.t9 * -33
		dp += tm.t10 * 110
		dp += tm.t11 * 32
		dp += tm.t12 * -56
		dp += tm.t13 * 110
		dp += tm.t14 * -67
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm12(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 32
		dp += tm.t1 * 18
		dp += tm.t2 * 27
		dp += tm.t3 * 17
		dp += tm.t4 * -14
		dp += tm.t5 * -28
		dp += tm.t6 * -41
		dp += tm.t7 * -15
		dp += tm.t8 * -4
		dp += tm.t9 * 147
		dp += tm.t10 * 37
		dp += tm.t11 * 125
		dp += tm.t12 * -56
		dp += tm.t13 * -97
		dp += tm.t14 * -67
		dp += tm.t15 * -31
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm13(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * -97
		dp += tm.t2 * -35
		dp += tm.t3 * -38
		dp += tm.t4 * -14
		dp += tm.t5 * 34
		dp += tm.t6 * -13
		dp += tm.t7 * 32
		dp += tm.t8 * 129
		dp += tm.t9 * -58
		dp += tm.t10 * -35
		dp += tm.t11 * 32
		dp += tm.t12 * -28
		dp += tm.t13 * -14
		dp += tm.t14 * 57
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(-166 * dp)
	}
}
func svm14(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * 18
		dp += tm.t2 * -3
		dp += tm.t3 * 17
		dp += tm.t4 * -14
		dp += tm.t5 * 3
		dp += tm.t6 * -68
		dp += tm.t7 * 32
		dp += tm.t8 * 102
		dp += tm.t9 * -84
		dp += tm.t10 * -11
		dp += tm.t11 * 63
		dp += tm.t12 * -83
		dp += tm.t13 * 27
		dp += tm.t14 * 132
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm15(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -34
		dp += tm.t1 * 0
		dp += tm.t2 * -67
		dp += tm.t3 * -10
		dp += tm.t4 * -44
		dp += tm.t5 * -123
		dp += tm.t6 * 13
		dp += tm.t7 * 32
		dp += tm.t8 * 156
		dp += tm.t9 * -33
		dp += tm.t10 * -35
		dp += tm.t11 * 186
		dp += tm.t12 * -83
		dp += tm.t13 * -14
		dp += tm.t14 * 82
		dp += tm.t15 * 7
		dp = dp*2
		res <- uint32(-665 * dp)
	}
}
func svm16(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * -78
		dp += tm.t2 * -3
		dp += tm.t3 * -95
		dp += tm.t4 * -44
		dp += tm.t5 * -60
		dp += tm.t6 * 13
		dp += tm.t7 * 8
		dp += tm.t8 * 48
		dp += tm.t9 * 44
		dp += tm.t10 * 86
		dp += tm.t11 * 32
		dp += tm.t12 * -1
		dp += tm.t13 * 27
		dp += tm.t14 * -42
		dp += tm.t15 * -71
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm17(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * 18
		dp += tm.t2 * -35
		dp += tm.t3 * -10
		dp += tm.t4 * -73
		dp += tm.t5 * -186
		dp += tm.t6 * 151
		dp += tm.t7 * 8
		dp += tm.t8 * -4
		dp += tm.t9 * 95
		dp += tm.t10 * 110
		dp += tm.t11 * -28
		dp += tm.t12 * -83
		dp += tm.t13 * -14
		dp += tm.t14 * -42
		dp += tm.t15 * -71
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm18(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * 0
		dp += tm.t1 * 38
		dp += tm.t2 * -35
		dp += tm.t3 * 17
		dp += tm.t4 * -14
		dp += tm.t5 * -28
		dp += tm.t6 * -13
		dp += tm.t7 * 102
		dp += tm.t8 * 48
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * -28
		dp += tm.t12 * -28
		dp += tm.t13 * -14
		dp += tm.t14 * 132
		dp += tm.t15 * 86
		dp = dp*2
		res <- uint32(33 * dp)
	}
}
func svm19(temp <-chan data, res chan<- uint32){
	for{
		var dp fixed.Int26_6 = 0
		tm := <-temp
		dp += tm.t0 * -34
		dp += tm.t1 * -58
		dp += tm.t2 * -35
		dp += tm.t3 * 17
		dp += tm.t4 * -44
		dp += tm.t5 * -123
		dp += tm.t6 * 13
		dp += tm.t7 * 32
		dp += tm.t8 * 129
		dp += tm.t9 * -33
		dp += tm.t10 * -11
		dp += tm.t11 * 186
		dp += tm.t12 * -83
		dp += tm.t13 * -14
		dp += tm.t14 * 82
		dp += tm.t15 * -31
		dp = dp*2
		res <- uint32(-99 * dp)
	}
}
func input0(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input1(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input2(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input3(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input4(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input5(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input6(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input7(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input8(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input9(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input10(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input11(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input12(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input13(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input14(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input15(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input16(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input17(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input18(in <-chan data, out chan<- data, next chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
		next <- tm
	}
}
func input19(in <-chan data, out chan<- data){
	var tm data
	for{
		tm = <- in
		out <- tm
	}
}
func output0(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output1(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output2(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output3(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output4(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output5(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output6(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output7(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output8(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output9(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output10(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output11(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output12(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output13(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output14(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output15(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output16(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output17(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}
func output18(op1chan <-chan uint32, op2chan <-chan uint32, out chan<- uint32){
	for{
		op1 := <- op1chan
		op2 := <- op2chan
		out <- op1+op2
	}
}

func Top(
	// Three operands from the host. Pointers to the input data and the space for the result in shared
	// memory and the length of the input data so the FPGA knows what to expect.
	input uintptr,
	outputData uintptr,
	//sv uintptr,
	//alpha uintptr,
	length uint32,
	// Set up channels for interacting with the shared memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	in0 := make(chan data, 2)
	in1 := make(chan data, 2)
	in2 := make(chan data, 2)
	in3 := make(chan data, 2)
	in4 := make(chan data, 2)
	in5 := make(chan data, 2)
	in6 := make(chan data, 2)
	in7 := make(chan data, 2)
	in8 := make(chan data, 2)
	in9 := make(chan data, 2)
	in10 := make(chan data, 2)
	in11 := make(chan data, 2)
	in12 := make(chan data, 2)
	in13 := make(chan data, 2)
	in14 := make(chan data, 2)
	in15 := make(chan data, 2)
	in16 := make(chan data, 2)
	in17 := make(chan data, 2)
	in18 := make(chan data, 2)
	in19 := make(chan data, 2)

	out0 := make(chan data, 2)
	out1 := make(chan data, 2)
	out2 := make(chan data, 2)
	out3 := make(chan data, 2)
	out4 := make(chan data, 2)
	out5 := make(chan data, 2)
	out6 := make(chan data, 2)
	out7 := make(chan data, 2)
	out8 := make(chan data, 2)
	out9 := make(chan data, 2)
	out10 := make(chan data, 2)
	out11 := make(chan data, 2)
	out12 := make(chan data, 2)
	out13 := make(chan data, 2)
	out14 := make(chan data, 2)
	out15 := make(chan data, 2)
	out16 := make(chan data, 2)
	out17 := make(chan data, 2)
	out18 := make(chan data, 2)
	out19 := make(chan data, 2)

	res0 := make(chan uint32, 2)
	res1 := make(chan uint32, 2)
	res2 := make(chan uint32, 2)
	res3 := make(chan uint32, 2)
	res4 := make(chan uint32, 2)
	res5 := make(chan uint32, 2)
	res6 := make(chan uint32, 2)
	res7 := make(chan uint32, 2)
	res8 := make(chan uint32, 2)
	res9 := make(chan uint32, 2)
	res10 := make(chan uint32, 2)
	res11 := make(chan uint32, 2)
	res12 := make(chan uint32, 2)
	res13 := make(chan uint32, 2)
	res14 := make(chan uint32, 2)
	res15 := make(chan uint32, 2)
	res16 := make(chan uint32, 2)
	res17 := make(chan uint32, 2)
	res18 := make(chan uint32, 2)
	res19 := make(chan uint32, 2)

	sum0 := make(chan uint32, 2)
	sum1 := make(chan uint32, 2)
	sum2 := make(chan uint32, 2)
	sum3 := make(chan uint32, 2)
	sum4 := make(chan uint32, 2)
	sum5 := make(chan uint32, 2)
	sum6 := make(chan uint32, 2)
	sum7 := make(chan uint32, 2)
	sum8 := make(chan uint32, 2)
	sum9 := make(chan uint32, 2)
	sum10 := make(chan uint32, 2)
	sum11 := make(chan uint32, 2)
	sum12 := make(chan uint32, 2)
	sum13 := make(chan uint32, 2)
	sum14 := make(chan uint32, 2)
	sum15 := make(chan uint32, 2)
	sum16 := make(chan uint32, 2)
	sum17 := make(chan uint32, 2)
	sum18 := make(chan uint32, 2)

	inputChan := make(chan uint32, 16)
	go aximemory.ReadBurstUInt32(
	memReadAddr, memReadData, true, input, length*16, inputChan)

	go func(){
		for i:=0; i<int(length);i++{
			t0 := fixed.Int26_6(<-inputChan)
			t1 := fixed.Int26_6(<-inputChan)
			t2 := fixed.Int26_6(<-inputChan)
			t3 := fixed.Int26_6(<-inputChan)
			t4 := fixed.Int26_6(<-inputChan)
			t5 := fixed.Int26_6(<-inputChan)
			t6 := fixed.Int26_6(<-inputChan)
			t7 := fixed.Int26_6(<-inputChan)
			t8 := fixed.Int26_6(<-inputChan)
			t9 := fixed.Int26_6(<-inputChan)
			t10 := fixed.Int26_6(<-inputChan)
			t11 := fixed.Int26_6(<-inputChan)
			t12 := fixed.Int26_6(<-inputChan)
			t13 := fixed.Int26_6(<-inputChan)
			t14 := fixed.Int26_6(<-inputChan)
			t15 := fixed.Int26_6(<-inputChan)
			temp := data{t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15}
			in0 <- temp
		}
	}()

	 go input0(in0, out0, in1)
	 go input1(in1, out1, in2)
	 go input2(in2, out2, in3)
	 go input3(in3, out3, in4)
	 go input4(in4, out4, in5)
	 go input5(in5, out5, in6)
	 go input6(in6, out6, in7)
	 go input7(in7, out7, in8)
	 go input8(in8, out8, in9)
	 go input9(in9, out9, in10)
	 go input10(in10, out10, in11)
	 go input11(in11, out11, in12)
	 go input12(in12, out12, in13)
	 go input13(in13, out13, in14)
	 go input14(in14, out14, in15)
	 go input15(in15, out15, in16)
	 go input16(in16, out16, in17)
	 go input17(in17, out17, in18)
	 go input18(in18, out18, in19)
	 go input19(in19, out19)

	 go svm0(out0, res0)
	 go svm1(out1, res1)
	 go svm2(out2, res2)
	 go svm3(out3, res3)
	 go svm4(out4, res4)
	 go svm5(out5, res5)
	 go svm6(out6, res6)
	 go svm7(out7, res7)
	 go svm8(out8, res8)
	 go svm9(out9, res9)
	 go svm10(out10, res10)
	 go svm11(out11, res11)
	 go svm12(out12, res12)
	 go svm13(out13, res13)
	 go svm14(out14, res14)
	 go svm15(out15, res15)
	 go svm16(out16, res16)
	 go svm17(out17, res17)
	 go svm18(out18, res18)
	 go svm19(out19, res19)

	 go output0(res0, res1, sum0)
	 go output1(res2, sum0, sum1)
	 go output2(res3, sum1, sum2)
	 go output3(res4, sum2, sum3)
	 go output4(res5, sum3, sum4)
	 go output5(res6, sum4, sum5)
	 go output6(res7, sum5, sum6)
	 go output7(res8, sum6, sum7)
	 go output8(res9, sum7, sum8)
	 go output9(res10, sum8, sum9)
	 go output10(res11, sum9, sum10)
	 go output11(res12, sum10, sum11)
	 go output12(res13, sum11, sum12)
	 go output13(res14, sum12, sum13)
	 go output14(res15, sum13, sum14)
	 go output15(res16, sum14, sum15)
	 go output16(res17, sum15, sum16)
	 go output17(res18, sum16, sum17)
	 go output18(res19, sum17, sum18)
		
	aximemory.WriteBurstUInt32(memWriteAddr, memWriteData, memWriteResp, true, outputData, length, sum18)
	// Once that's done, we can exit.
}

