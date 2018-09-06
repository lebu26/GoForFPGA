package main
//Non-for loop version
import (
	// Import the entire framework for interracting with SDAccel from Go (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	
  // Use the new AXI protocol package for interracting with memory
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
	//FIR "github.com/lebu26/fir_100"
)

func Top(
	// Three operands from the host. Pointers to the input data and the space for the result in shared
	// memory and the length of the input data so the FPGA knows what to expect.
	xAddr uintptr,
	yAddr uintptr,
	length uint32,
	// Set up channels for interacting with the shared memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	x := make(chan uint32, 2)

	go aximemory.ReadBurstUInt32(
		memReadAddr, memReadData, true, xAddr, length, x)

	aximemory.WriteBurstUInt32(
		memWriteAddr, memWriteData, memWriteResp, true, yAddr, length, x)
}
