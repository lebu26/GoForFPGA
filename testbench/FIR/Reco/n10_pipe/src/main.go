package main
//Non-for loop version
import (
	// Import the entire framework for interracting with SDAccel from Go (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	
  // Use the new AXI protocol package for interracting with memory
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)


func acc1(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}
func acc2(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}
func acc3(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}
func acc4(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}
func acc5(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}
func acc6(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}
func acc7(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}
func acc8(y1 <-chan int, y2 <-chan int, out chan<-int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 +op2
	}
}

func fir1(x1 <-chan int,  y1 chan<- int){
	for{
		temp := int(<-x1)
		y1 <- temp * 1 
	}	
}

func fir2(x2 <-chan int, y2 chan<- int){
	for{
		temp := <-x2
		y2 <- temp * 2
	}	
}
func fir3(x3 <-chan int, y3 chan<- int){
	for{
		temp := <-x3
		y3 <- temp * 3
	}	
}
func fir4(x4 <-chan int, y4 chan<- int){
	for{
		temp := <-x4
		y4 <- temp * 4
	}	
}
func fir5(x5 <-chan int, y5 chan<- int){
	for{
		temp := <-x5
		y5 <- temp * 5
	}	
}
func fir6(x6 <-chan int,  y6 chan<- int){
	for{
		temp := int(<-x6)
		y6 <- temp * 6 
	}	
}

func fir7(x7 <-chan int, y7 chan<- int){
	for{
		temp := <-x7
		y7 <- temp * 7
	}	
}
func fir8(x8 <-chan int, y8 chan<- int){
	for{
		temp := <-x8
		y8 <- temp * 8
	}	
}
func fir9(x9 <-chan int, y9 chan<- int){
	for{
		temp := <-x9
		y9 <- temp * 9
	}	
}
func fir10(x10 <-chan int, y10 chan<- int){
	for{
		temp := <-x10
		y10 <- temp * 10
	}	
}


func input_first(in <-chan uint32, out chan<- int, move chan<- int){
	var x int = 0
	for{
		x = int(<-in)
		out <- x
		move <- x
	}
}

func input1(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}
func input2(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}
func input3(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}
func input4(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}
func input5(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}
func input6(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}
func input7(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}
func input8(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <- x
		move <- x
		x = <-in
	}
}

func input_last(in <-chan int, out chan<- int){
	var x int = 0
	for{
		out <- x
		x = <-in
	}
}

func out_control(y10 <-chan int, out8 <-chan int, y chan<- uint32){
	for{
		op1 := <- y10
		op2 := <- out8
		y <- uint32(op1+op2)
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

	x1 := make(chan int, 2)
	y1 := make(chan int, 2)

	x2 := make(chan int, 2)
	y2 := make(chan int, 2)

	x3 := make(chan int, 2)
	y3 := make(chan int, 2)

	x4 := make(chan int, 2)
	y4 := make(chan int, 2)

	x5 := make(chan int, 2)
	y5 := make(chan int, 2)

	x6 := make(chan int, 2)
	y6 := make(chan int, 2)

	x7 := make(chan int, 2)
	y7 := make(chan int, 2)

	x8 := make(chan int, 2)
	y8 := make(chan int, 2)

	x9 := make(chan int, 2)
	y9 := make(chan int, 2)

	x10 := make(chan int, 2)
	y10 := make(chan int, 2)

	y := make(chan uint32, 2)
	x := make(chan uint32, 2)

	out1 := make(chan int,2)
	out2 := make(chan int,2)	
	out3 := make(chan int,2)
	out4 := make(chan int,2)	
	out5 := make(chan int,2)
	out6 := make(chan int,2)
	out7 := make(chan int,2)
	out8 := make(chan int,2)	

	in1 := make(chan int,2)
	in2 := make(chan int,2)	
	in3 := make(chan int,2)
	in4 := make(chan int,2)	
	in5 := make(chan int,2)
	in6 := make(chan int,2)
	in7 := make(chan int,2)
	in8 := make(chan int,2)	
	in9 := make(chan int,2)

	go aximemory.ReadBurstUInt32(
		memReadAddr, memReadData, true, xAddr, length, x)

	go fir1(x1,y1)
	go fir2(x2,y2)
	go fir3(x3,y3)
	go fir4(x4,y4)
	go fir5(x5,y5)
	go fir6(x6,y6)
	go fir7(x7,y7)
	go fir8(x8,y8)
	go fir9(x9,y9)
	go fir10(x10,y10)
	//go acc(y1,y2,y3,y4,y5,y6,y7,y8,y9,y10, y)
	go acc1(y1, y2, out1)
	go acc2(y3,out1,out2)
	go acc3(y4,out2,out3)
	go acc4(y5,out3,out4)
	go acc5(y6,out4,out5)
	go acc6(y7,out5,out6)
	go acc7(y8,out6,out7)
	go acc8(y9,out7,out8)
	go out_control(y10,out8,y)

	go input_first(x,x1,in1)
	go input1(in1,x2,in2)
	go input2(in2,x3,in3)
	go input3(in3,x4,in4)
	go input4(in4,x5,in5)
	go input5(in5,x6,in6)
	go input6(in6,x7,in7)
	go input7(in7,x8,in8)
	go input8(in8,x9,in9)
	go input_last(in9,x10)
	//go input_control(x, x1,x2,x3,x4,x5,x6,x7,x8,x9,x10)
/*
	done := make(chan int)
	go func(){
			var res uint32
			for w:=0;w<int(length);w++{
				res = uint32(<- y)
				aximemory.WriteUInt32(memWriteAddr, memWriteData, memWriteResp, true,
				yAddr+uintptr(w*4),res)
			}
			done <- 1
		}()
	<-done
*/	
	aximemory.WriteBurstUInt32(
		memWriteAddr, memWriteData, memWriteResp, true, yAddr, length, y)
}
