package main

import (
	//"os"
	//"strconv"
	"encoding/binary"
	"time"
	"github.com/ReconfigureIO/sdaccel/xcl"
	"log"
	//"github.com/lebu26/testdata1"
	"math/rand"
)


func main(){
	const N int = 30 

	//var length uint32 = 1 
	length :=[2]int{10000000,100000000}

	log.Printf("**********************program starting*********************\n\r")
	log.Printf("The number of test cases is %v \n\r", length[0]);
	log.Printf("Software part staring...\n\r")
	
	testdata := make([]int32, length[0])
	gold:= make([]int32, length[0])
	//coef1 := [N]int32{2, 8, 5, 1, 10, 5, 9, 9, 3, 5, 6, 6, 2, 8, 2, 2, 6, 3, 8, 7, 2, 5, 3, 4, 3, 3, 2, 7, 9, 6, 8, 7, 2, 9, 10, 3, 8, 10, 6, 5, 4, 2, 3, 4, 4, 5, 2, 2, 4, 9, 8, 5, 3, 8, 8, 10, 4, 2, 10, 9, 7, 6, 1, 3, 9, 7, 1, 3, 5, 9, 7, 6, 1, 10, 1, 1, 7, 2, 4, 9, 10, 4, 5, 5, 7, 1, 7, 7, 2, 9, 5, 10, 7, 4, 8, 9, 9, 3, 10, 2}
	coef := [N]int32{2, 8, 5, 1, 10, 5, 9, 9, 3, 5, 6, 6, 2, 8, 2, 2, 6, 3, 8, 7, 2, 5, 3, 4, 3, 3, 2, 7, 9, 6}
	X := [N]int32{0}
	var temp int32 =0
	t1_s := time.Now()
	for i:=0;i<length[0];i++{
		testdata[i] = int32(rand.Intn(8) - 4)
		X[N-1] = testdata[i]
		temp = 0
		for j:=1;j<=N;j++{
			temp += X[N-j]*coef[j-1]
		}
		gold[i] = temp
		for w:=0;w<N-1;w++{
			X[w] = X[w+1]
		}
	}
	
	t2_sw := time.Now()
	td_sw := t2_sw.Sub(t1_s)
	log.Printf("**********************software finished*********************\n\r\r\r")
	log.Printf("** takes = %v s\n\r", td_sw.Seconds())

	log.Printf("************************************************************************************* \n\r")
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
	output := make([]uint32, length[0])

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


	// Pass the length[0] of the vector as the fourth argument
	krnl.SetArg(2, uint32(length[0]))
	log.Printf("kernel arguments initialization finished...\n\r")

	// Run the FPGA with the supplied arguments. This is the same for all projects.
	// The arguments ``(1, 1, 1)`` relate to x, y, z co-ordinates and correspond to our current
	// underlying technology.
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel staring\n\r")
	log.Printf("************************************************************************************* \n\r")
	t1_hw := time.Now()
	krnl.Run(1, 1, 1)

	// Read the result from shared memory. If it is zero return an error
	err := binary.Read(outputBuff.Reader(), binary.LittleEndian, output)
	if err != nil {
		log.Fatal("binary.Read failed:", err)
	}
	t2_hw := time.Now()
	td_hw = t2_hw.Sub(t1_hw)

	err_count := 0
	for w:=0 ; w<length[0]; w++{
		if int32(output[w]) != gold[w]{
			err_count ++
			if err_count<10{
				log.Printf("** output[%d] = %v, gold[%d] = %v \n\r", w,int32(output[w]),w,gold[w])
			}
		}
		

	} 

	log.Printf("**********************hardware finished*********************\n\r\r\r")
	log.Printf("** takes = %v s\n\r", td_hw.Seconds())
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel finished successfully\n\r")
	log.Printf("                                  error count = %v \n\r",err_count)
	log.Printf("************************************************************************************* \n\r")
///*************************************************************************************************************************************************
//SECOND TEST



	log.Printf("**********************program starting*********************\n\r")
	log.Printf("The number of test cases is %v \n\r", length[1]);
	log.Printf("Software part staring...\n\r")
	t1_sw1 := time.Now()
	testdata1 := make([]int32, length[1])
	gold1 := make([]int32, length[1])
	//coef := [N]int32{2, 8, 5, 1, 10, 5, 9, 9, 3, 5, 6, 6, 2, 8, 2, 2, 6, 3, 8, 7, 2, 5, 3, 4, 3, 3, 2, 7, 9, 6, 8, 7, 2, 9, 10, 3, 8, 10, 6, 5, 4, 2, 3, 4, 4, 5, 2, 2, 4, 9, 8, 5, 3, 8, 8, 10, 4, 2, 10, 9, 7, 6, 1, 3, 9, 7, 1, 3, 5, 9, 7, 6, 1, 10, 1, 1, 7, 2, 4, 9, 10, 4, 5, 5, 7, 1, 7, 7, 2, 9, 5, 10, 7, 4, 8, 9, 9, 3, 10, 2}
	coef1 := [N]int32{2, 8, 5, 1, 10, 5, 9, 9, 3, 5, 6, 6, 2, 8, 2, 2, 6, 3, 8, 7, 2, 5, 3, 4, 3, 3, 2, 7, 9, 6}
	X1 := [N]int32{0}
	var temp1 int32 =0
	for i:=0;i<length[1];i++{
		testdata1[i] = int32(rand.Intn(8) - 4)
		X1[N-1] = testdata1[i]
		temp1 = 0
		for j:=1;j<=N;j++{
			temp1 += X1[N-j]*coef1[j-1]
		}
		gold1[i] = temp1
		for w:=0;w<N-1;w++{
			X1[w] = X1[w+1]
		}
	}
	
	t2_sw1 := time.Now()
	td_sw1 := t2_sw1.Sub(t1_sw1)
	log.Printf("**********************software finished*********************\n\r\r\r")
	log.Printf("** takes = %v s\n\r", td_sw1.Seconds())

	log.Printf("************************************************************************************* \n\r")
	log.Printf("hardware part staring...\n\r")


	// Allocate a 'world' for interacting with the FPGA
	world1 := xcl.NewWorld()
	defer world1.Release()

	// Import the compiled code that will be loaded onto the FPGA (referred to here as a kernel)
	// Right now these two identifiers are hard coded as an output from the build process
	krnl1 := world1.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl1.Release()

	// Allocate a space in the shared memory to store the data you're sending to the FPGA
	buff_in1 := world1.Malloc(xcl.ReadOnly, uint(binary.Size(testdata1)))
	defer buff_in1.Free()

	// Construct an array to hold the output data from the FPGA
	output1 := make([]uint32, length[1])

	// Allocate a space in the shared memory to store the output1 data from the FPGA.
	output1Buff := world1.Malloc(xcl.ReadWrite, uint(binary.Size(output1)))
	defer output1Buff.Free()

	// Write our input data to shared memory at the address we previously allocated
	binary.Write(buff_in1.Writer(), binary.LittleEndian, testdata1)

	// Zero out the space in shared memory for the result from the FPGA
	binary.Write(output1Buff.Writer(), binary.LittleEndian, output1)

	// Pass the pointer to the input data in shared memory as the first and second argument
	krnl1.SetMemoryArg(0, buff_in1)
	//Pass the pointer to the output1 data in shared memory as the third argument
	krnl1.SetMemoryArg(1, output1Buff)

	var td_hw1 time.Duration


	// Pass the length[1] of the vector as the fourth argument
	krnl1.SetArg(2, uint32(length[1]))
	log.Printf("kernel arguments initialization finished...\n\r")

	// Run the FPGA with the supplied arguments. This is the same for all projects.
	// The arguments ``(1, 1, 1)`` relate to x, y, z co-ordinates and correspond to our current
	// underlying technology.
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel staring\n\r")
	log.Printf("************************************************************************************* \n\r")
	t1_hw1 := time.Now()
	krnl1.Run(1, 1, 1)

	// Read the result from shared memory. If it is zero return an error
	err = binary.Read(output1Buff.Reader(), binary.LittleEndian, output1)
	if err != nil {
		log.Fatal("binary.Read failed:", err)
	}
	t2_hw1 := time.Now()
	td_hw1 = t2_hw1.Sub(t1_hw1)

	err_count = 0
	for w:=0 ; w<length[1]; w++{
		if int32(output1[w]) != gold1[w]{
			err_count ++
			if err_count<10{
				log.Printf("** output[%d] = %v, gold1[%d] = %v \n\r", w,int32(output1[w]),w,gold1[w])
			}
		}
		

	} 

	log.Printf("**********************hardware finished*********************\n\r\r\r")
	log.Printf("** takes = %v s\n\r", td_hw1.Seconds())
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel finished successfully\n\r")
	log.Printf("                                  error count = %v \n\r",err_count)
	log.Printf("                                  for %v cases: \n\r", length[0])
	log.Printf(" ---------------------------------hardware takes = %v s\n\r", td_hw.Seconds())
	log.Printf(" ---------------------------------software takes = %v s\n\r", td_sw.Seconds())
	log.Printf("                                  for %v cases: \n\r", length[1])
	log.Printf(" ---------------------------------hardware takes = %v s\n\r", td_hw1.Seconds())
	log.Printf(" ---------------------------------software takes = %v s\n\r", td_sw1.Seconds())
	log.Printf("************************************************************************************* \n\r")

}