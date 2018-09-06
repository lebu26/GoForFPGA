package main

import (
	// Import the entire framework for interracting with SDAccel from Go (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	
  // Use the new AXI protocol package for interracting with memory
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)

const N int = 10

func FIR(input <-chan uint32, output chan<- uint32){
	coef := [N]int{1,2,3,4,5,6,7,8,9,10}
	X := [N]int{0,0,0,0,0,0,0,0,0,0}
	var temp int = 0
	for {
		X[N-1] = int(<-input)
		temp =0
		for j:=1;j<=N;j++{
			temp += X[N-j]*coef[j-1]
		}
		output <- uint32(temp)
		for w:=0;w<N-1;w++{
			X[w] = X[w+1]
		}
	}
}

func Top(
	// Three operands from the host. Pointers to the input data and the space for the result in shared
	// memory and the length of the input data so the FPGA knows what to expect.
	xAddr uintptr,
	yAddr uintptr,
	//sv uintptr,
	//alpha uintptr,
	length uint32,
	// Set up channels for interacting with the shared memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

/* old version

	input := make(chan uint32)
	go aximemory.ReadBurstUInt32(
		memReadAddr, memReadData, true, xAddr, length, input)

	coef := [N]int{1,2,3,4,5,6,7,8,9,10}
	X := [N]int{0,0,0,0,0,0,0,0,0,0}
	var temp int = 0
	for i:=0;i<int(length);i++{
		temp =0
		X[N-1] = int(<-input)
		for j:=1;j<=N;j++{
			temp += X[N-j]*coef[j-1]
		}
		aximemory.WriteUInt32(
			memWriteAddr, memWriteData, memWriteResp, true,
			yAddr+uintptr(i*4),uint32(temp))
		for w:=1;w<N;w++{
			X[N-1-w] = X[N-w]
		}
	}
*/
	input := make(chan uint32)
	output := make(chan uint32)

	go aximemory.ReadBurstUInt32(
		memReadAddr, memReadData, true, xAddr, length, input)

	go FIR(input, output)
	
	aximemory.WriteBurstUInt32(
		memWriteAddr, memWriteData, memWriteResp, true, yAddr, length, output)
}
