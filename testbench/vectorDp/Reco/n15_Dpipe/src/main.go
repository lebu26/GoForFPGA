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

func svm_loop0(temp <-chan data, sum0_OC chan<- uint32){ 
	var tm data
	for{
		tm = <-temp 
		var sum fixed.Int26_6 = 0 
		var dp fixed.Int26_6 = 0 
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += -832 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += -832 * dp
		dp = 0
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
		sum += -832 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += 33 * dp
		dp = 0
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
		sum += -166 * dp
		dp = 0
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
		sum += 33 * dp
		
		sum0_OC <- uint32(sum)
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

	inputChan := make(chan uint32)
	/*
	for i:=0; i<int(length*16);i++{
		
		inputChan<- aximemory.ReadUInt32(
		memReadAddr, memReadData, true, sv+uintptr(4*i))
		
	}
	*/	
	svm_loop0chan := make(chan uint32, 2)
	svm_input0chan := make(chan data, 2)

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
			svm_input0chan <- temp
		}
	}()

	go svm_loop0(svm_input0chan, svm_loop0chan)
		
	aximemory.WriteBurstUInt32(memWriteAddr, memWriteData, memWriteResp, true, outputData, length, svm_loop0chan)
	// Once that's done, we can exit.
}
