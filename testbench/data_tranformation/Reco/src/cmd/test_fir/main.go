package main

import (
	"os"
	"strconv"
	"encoding/binary"
	"time"
	"github.com/ReconfigureIO/sdaccel/xcl"
	"log"
	//"github.com/lebu26/TestData"
	"math/rand"
)


func main(){
	const N int = 20 

	//var length uint32 = 1 
	arg := os.Args[1]
	length, error_t := strconv.Atoi(arg)

	if error_t == nil{

	}
	log.Printf("**********************program starting*********************\n\r")
	log.Printf("The arg is %s \n\r", arg);
	log.Printf("The number of test cases is %v \n\r", length);

	testdata := make([]int32, length)

	for i:=0;i<length;i++{
		testdata[i] = int32(rand.Intn(8) - 4)
	}

	log.Printf("hardware part staring...\n\r")

	// Allocate a 'world' for interacting with the FPGA
	world := xcl.NewWorld()
	defer world.Release()

	// Import the compiled code that will be loaded onto the FPGA (referred to here as a kernel)
	// Right now these two identifiers are hard coded as an output from the build process
	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	// Allocate a space in the shared memory to store the data you're sending to the FPGA
	buff_in := world.Malloc(xcl.ReadOnly, uint(binary.Size(testdata)))
	defer buff_in.Free()

	// Construct an array to hold the output data from the FPGA
	output := make([]uint32, length)

	// Allocate a space in the shared memory to store the output data from the FPGA.
	outputBuff := world.Malloc(xcl.ReadWrite, uint(binary.Size(output)))
	defer outputBuff.Free()

	// Write our input data to shared memory at the address we previously allocated
	binary.Write(buff_in.Writer(), binary.LittleEndian, testdata)

	// Zero out the space in shared memory for the result from the FPGA
	binary.Write(outputBuff.Writer(), binary.LittleEndian, output)

	// Pass the pointer to the input data in shared memory as the first and second argument
	krnl.SetMemoryArg(0, buff_in)
	//Pass the pointer to the output data in shared memory as the third argument
	krnl.SetMemoryArg(1, outputBuff)

	var td_hw time.Duration


	// Pass the length of the vector as the fourth argument
	krnl.SetArg(2, uint32(length))
	log.Printf("kernel arguments initialization finished...\n\r")

	// Run the FPGA with the supplied arguments. This is the same for all projects.
	// The arguments ``(1, 1, 1)`` relate to x, y, z co-ordinates and correspond to our current
	// underlying technology.
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel staring\n\r")
	log.Printf("************************************************************************************* \n\r")
	var t1_hw = time.Now()
	krnl.Run(1, 1, 1)

	// Read the result from shared memory. If it is zero return an error
	err := binary.Read(outputBuff.Reader(), binary.LittleEndian, output)
	if err != nil {
		log.Fatal("binary.Read failed:", err)
	}
	var t2_hw = time.Now()
	td_hw = t2_hw.Sub(t1_hw)
	log.Printf("**********************hardware finished*********************\n\r\r\r")
	log.Printf("** takes = %v s\n\r", td_hw.Seconds())
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel finished successfully\n\r")
	log.Printf("************************************************************************************* \n\r")
}