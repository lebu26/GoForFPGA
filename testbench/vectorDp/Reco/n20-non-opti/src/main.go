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

func svm(temp <-chan data, res chan<- uint32) {
		SV := [320]fixed.Int26_6{0, -97, -3, -38, 14, 34, -13, 32, 22, -33, -11, -59, -28, -14, 82, 86, -34, -97, -67, -67, -44, 34, 13, 55, -4, -33, -35, -28, -28, -14, 132, 7, -34, -58, -67, 17, -44, -123, 13, 32, 129, -33, -11, 217, -83, -14, 82, 7, -34, 38, -35, 45, -44, 3, -68, 32, 129, -84, -11, 63, -83, 27, 132, 7, -67, -58, -35, -67, -44, -28, 68, -85, 48, -7, 110, 2, -56, 110, -42, -31, -34, 18, -67, 17, -44, -123, -13, 32, 129, -33, -11, 186, -83, -14, 82, 7, -101, -136, -131, -123, -73, -91, -13, 8, 75, -33, -11, 156, -83, -14, 57, 86, -34, -20, -35, -38, -14, 3, 123, -15, 22, -33, 110, 2, -28, 152, -67, -31, 0, -97, -3, -38, 14, 3, -13, 8, 22, -33, -11, -59, -28, -14, 82, 86, 32, -39, -3, 74, 14, 3, 41, 126, 75, -33, -35, -28, -28, -14, 132, 86, 0, 0, -35, -10, -14, -28, -13, 102, 48, -33, -11, -28, -28, -14, 132, 86, -67, -97, -35, -38, -44, -28, 68, -85, 48, -33, 110, 32, -56, 110, -67, 7, 32, 18, 27, 17, -14, -28, -41, -15, -4, 147, 37, 125, -56, -97, -67, -31, 0, -97, -35, -38, -14, 34, -13, 32, 129, -58, -35, 32, -28, -14, 57, 7, 0, 18, -3, 17, -14, 3, -68, 32, 102, -84, -11, 63, -83, 27, 132, 7, -34, 0, -67, -10, -44, -123, 13, 32, 156, -33, -35, 186, -83, -14, 82, 7, 0, -78, -3, -95, -44, -60, 13, 8, 48, 44, 86, 32, -1, 27, -42, -71, 0, 18, -35, -10, -73, -186, 151, 8, -4, 95, 110, -28, -83, -14, -42, -71, 0, 38, -35, 17, -14, -28, -13, 102, 48, -33, -11, -28, -28, -14, 132, 86, -34, -58, -35, 17, -44, -123, 13, 32, 129, -33, -11, 186, -83, -14, 82, -31, }
		Alpha := [20]fixed.Int26_6{33, 33, -832, 33, 33, -832, -832, 33, 33, 33, 33, 33, 33, -166, 33, -665, 33, 33, 33, -99, }
	for{
		tm :=  <-temp
		var sum fixed.Int26_6 = 0		

		for j := 0; j < 20;j++{
			var dp fixed.Int26_6 = 0
			dp += tm.t0*SV[j*16]
			dp += tm.t1*SV[j*16 + 1]
			dp += tm.t2*SV[j*16 + 2]
			dp += tm.t3*SV[j*16 + 3]
			dp += tm.t4*SV[j*16 + 4]
			dp += tm.t5*SV[j*16 + 5]
			dp += tm.t6*SV[j*16 + 6]
			dp += tm.t7*SV[j*16 + 7]
			dp += tm.t8*SV[j*16 + 8] 
			dp += tm.t9*SV[j*16 + 9]
			dp += tm.t10*SV[j*16 + 10]
			dp += tm.t11*SV[j*16 + 11]
			dp += tm.t12*SV[j*16 + 12]
			dp += tm.t13*SV[j*16 + 13]
			dp += tm.t14*SV[j*16 +14]
			dp += tm.t15*SV[j*16 +15]
			dp = dp *2
			sum += Alpha[j] *dp		
		}
		res <- uint32(sum)
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

	go svm(svm_input0chan, svm_loop0chan)
		
	aximemory.WriteBurstUInt32(memWriteAddr, memWriteData, memWriteResp, true, outputData, length, svm_loop0chan)
	// Once that's done, we can exit.
}

