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
		dp = tm.t0 + 
		(tm.t1 * -97)+
		(tm.t2 * -3)+
		(tm.t3 * -38)+
		(tm.t4 * 14)+
		(tm.t5 * 34)+
		(tm.t6 * -13)+
		(tm.t7 * 32)+
		(tm.t8 * 22)+
		(tm.t9 * -33)+
		(tm.t10 * -11)+
		(tm.t11 * -59)+
		(tm.t12 * -28)+
		(tm.t13 * -14)+
		(tm.t14 * 82)+
		(tm.t15 * 86)
		dp = dp*2
		//sum += 
		sum0_OC <- uint32(33 * dp)
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
		tmp :=[16]uint32{}
		for {
			for i := 0; i < 16; i++ {
				tmp[i] = <-inputChan
			}
			svm_input0chan <- data{
				fixed.Int26_6(tmp[0]),
				fixed.Int26_6(tmp[1]),
				fixed.Int26_6(tmp[2]),
				fixed.Int26_6(tmp[3]),
				fixed.Int26_6(tmp[4]),
				fixed.Int26_6(tmp[5]),
				fixed.Int26_6(tmp[6]),
				fixed.Int26_6(tmp[7]),
				fixed.Int26_6(tmp[8]),
				fixed.Int26_6(tmp[9]),
				fixed.Int26_6(tmp[10]),
				fixed.Int26_6(tmp[11]),
				fixed.Int26_6(tmp[12]),
				fixed.Int26_6(tmp[13]),
				fixed.Int26_6(tmp[14]),
				fixed.Int26_6(tmp[15]),
			}
		}
	}()

	go svm_loop0(svm_input0chan, svm_loop0chan)
		
	aximemory.WriteBurstUInt32(memWriteAddr, memWriteData, memWriteResp, true, outputData, length, svm_loop0chan)
	// Once that's done, we can exit.
}
