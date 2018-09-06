
package main

import (
	// Import the entire framework for interracting with SDAccel from Go (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	
  // Use the new AXI protocol package for interracting with memory
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"

	//"github.com/ReconfigureIO/fixed"
)

type data struct {
	t0 int
	t1 int
	t2 int
	t3 int
	t4 int
	t5 int
	t6 int
	t7 int 
	t8 int
	t9 int
	t10 int
	t11 int 
	t12 int
	t13 int
	t14 int
	t15 int  
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

	inputChan := make(chan uint32,2)
	in0 := make(chan data,2)
	out := make(chan uint32,2) 
	go aximemory.ReadBurstUInt32(
		memReadAddr, memReadData, true, input, length*16, inputChan)

	//bias := fixed.I26F(231, 344509)

	go func(){
		for i:=0; i<int(length);i++{
			t0 := int(<-inputChan)
			t1 := int(<-inputChan)
			t2 := int(<-inputChan)
			t3 := int(<-inputChan)
			t4 := int(<-inputChan)
			t5 := int(<-inputChan)
			t6 := int(<-inputChan)
			t7 := int(<-inputChan)
			t8 := int(<-inputChan)
			t9 := int(<-inputChan)
			t10 := int(<-inputChan)
			t11 := int(<-inputChan)
			t12 := int(<-inputChan)
			t13 := int(<-inputChan)
			t14 := int(<-inputChan)
			t15 := int(<-inputChan)
			temp := data{t0,t1,t2,t3,t4,t5,t6,t7,t8,t9,t10,t11,t12,t13,t14,t15}
			in0 <- temp
		}
	}()

	go func(){
		for i:=0; i<int(length);i++{
			tm := <-in0
			out <- uint32(tm.t15)
		}
	}()

	
	aximemory.WriteBurstUInt32(
	memWriteAddr, memWriteData, memWriteResp, true, outputData, length, out)
	
}

