package FIR

func acc0(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc1(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc2(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc3(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc4(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc5(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc6(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc7(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc8(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc9(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc10(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc11(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc12(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc13(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc14(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc15(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc16(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func acc17(y1 <-chan int, y2 <-chan int, out chan<- int){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- op1 + op2
	}
}
func out_control(y1 <-chan int, y2 <-chan int, out chan<- uint32){
	for{
		op1 := <-y1
		op2 := <-y2
		out <- uint32(op1 + op2)
	}
}
func fir0(x0 <-chan int, y0 chan<- int){
	for{
		temp := int(<-x0)
		y0 <- temp*2
	}
}
func fir1(x1 <-chan int, y1 chan<- int){
	for{
		temp := int(<-x1)
		y1 <- temp*8
	}
}
func fir2(x2 <-chan int, y2 chan<- int){
	for{
		temp := int(<-x2)
		y2 <- temp*5
	}
}
func fir3(x3 <-chan int, y3 chan<- int){
	for{
		temp := int(<-x3)
		y3 <- temp*1
	}
}
func fir4(x4 <-chan int, y4 chan<- int){
	for{
		temp := int(<-x4)
		y4 <- temp*10
	}
}
func fir5(x5 <-chan int, y5 chan<- int){
	for{
		temp := int(<-x5)
		y5 <- temp*5
	}
}
func fir6(x6 <-chan int, y6 chan<- int){
	for{
		temp := int(<-x6)
		y6 <- temp*9
	}
}
func fir7(x7 <-chan int, y7 chan<- int){
	for{
		temp := int(<-x7)
		y7 <- temp*9
	}
}
func fir8(x8 <-chan int, y8 chan<- int){
	for{
		temp := int(<-x8)
		y8 <- temp*3
	}
}
func fir9(x9 <-chan int, y9 chan<- int){
	for{
		temp := int(<-x9)
		y9 <- temp*5
	}
}
func fir10(x10 <-chan int, y10 chan<- int){
	for{
		temp := int(<-x10)
		y10 <- temp*6
	}
}
func fir11(x11 <-chan int, y11 chan<- int){
	for{
		temp := int(<-x11)
		y11 <- temp*6
	}
}
func fir12(x12 <-chan int, y12 chan<- int){
	for{
		temp := int(<-x12)
		y12 <- temp*2
	}
}
func fir13(x13 <-chan int, y13 chan<- int){
	for{
		temp := int(<-x13)
		y13 <- temp*8
	}
}
func fir14(x14 <-chan int, y14 chan<- int){
	for{
		temp := int(<-x14)
		y14 <- temp*2
	}
}
func fir15(x15 <-chan int, y15 chan<- int){
	for{
		temp := int(<-x15)
		y15 <- temp*2
	}
}
func fir16(x16 <-chan int, y16 chan<- int){
	for{
		temp := int(<-x16)
		y16 <- temp*6
	}
}
func fir17(x17 <-chan int, y17 chan<- int){
	for{
		temp := int(<-x17)
		y17 <- temp*3
	}
}
func fir18(x18 <-chan int, y18 chan<- int){
	for{
		temp := int(<-x18)
		y18 <- temp*8
	}
}
func fir19(x19 <-chan int, y19 chan<- int){
	for{
		temp := int(<-x19)
		y19 <- temp*7
	}
}
func input_first(in <-chan uint32, out chan<- int, move chan<- int){
	var x int = 0
	for{
		x = int(<-in)
		out <-x
		move <- x
	}
}
func input0(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input1(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input2(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input3(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input4(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input5(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input6(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input7(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input8(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input9(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input10(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input11(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input12(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input13(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input14(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input15(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input16(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input17(in <-chan int, out chan<- int, move chan<- int){
	var x int = 0
	for{
		out <-x
		move <- x
		x = <-in
	}
}
func input_last(in <-chan int, out chan<- int){
	var x int = 0
	for{
		out <-x
		x = <-in
	}
}
func FIR_Top(x <-chan uint32, y chan<- uint32){
	x0 := make(chan int, 2)
	y0 := make(chan int, 2)

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

	x11 := make(chan int, 2)
	y11 := make(chan int, 2)

	x12 := make(chan int, 2)
	y12 := make(chan int, 2)

	x13 := make(chan int, 2)
	y13 := make(chan int, 2)

	x14 := make(chan int, 2)
	y14 := make(chan int, 2)

	x15 := make(chan int, 2)
	y15 := make(chan int, 2)

	x16 := make(chan int, 2)
	y16 := make(chan int, 2)

	x17 := make(chan int, 2)
	y17 := make(chan int, 2)

	x18 := make(chan int, 2)
	y18 := make(chan int, 2)

	x19 := make(chan int, 2)
	y19 := make(chan int, 2)

	out0 := make(chan int, 2)
	out1 := make(chan int, 2)
	out2 := make(chan int, 2)
	out3 := make(chan int, 2)
	out4 := make(chan int, 2)
	out5 := make(chan int, 2)
	out6 := make(chan int, 2)
	out7 := make(chan int, 2)
	out8 := make(chan int, 2)
	out9 := make(chan int, 2)
	out10 := make(chan int, 2)
	out11 := make(chan int, 2)
	out12 := make(chan int, 2)
	out13 := make(chan int, 2)
	out14 := make(chan int, 2)
	out15 := make(chan int, 2)
	out16 := make(chan int, 2)
	out17 := make(chan int, 2)

	in0 := make(chan int, 2)
	in1 := make(chan int, 2)
	in2 := make(chan int, 2)
	in3 := make(chan int, 2)
	in4 := make(chan int, 2)
	in5 := make(chan int, 2)
	in6 := make(chan int, 2)
	in7 := make(chan int, 2)
	in8 := make(chan int, 2)
	in9 := make(chan int, 2)
	in10 := make(chan int, 2)
	in11 := make(chan int, 2)
	in12 := make(chan int, 2)
	in13 := make(chan int, 2)
	in14 := make(chan int, 2)
	in15 := make(chan int, 2)
	in16 := make(chan int, 2)
	in17 := make(chan int, 2)
	in18 := make(chan int, 2)

	go fir0(x0, y0)
	go fir1(x1, y1)
	go fir2(x2, y2)
	go fir3(x3, y3)
	go fir4(x4, y4)
	go fir5(x5, y5)
	go fir6(x6, y6)
	go fir7(x7, y7)
	go fir8(x8, y8)
	go fir9(x9, y9)
	go fir10(x10, y10)
	go fir11(x11, y11)
	go fir12(x12, y12)
	go fir13(x13, y13)
	go fir14(x14, y14)
	go fir15(x15, y15)
	go fir16(x16, y16)
	go fir17(x17, y17)
	go fir18(x18, y18)
	go fir19(x19, y19)

	go acc0(y0, y1, out0)
	go acc1(y2, out0, out1)
	go acc2(y3, out1, out2)
	go acc3(y4, out2, out3)
	go acc4(y5, out3, out4)
	go acc5(y6, out4, out5)
	go acc6(y7, out5, out6)
	go acc7(y8, out6, out7)
	go acc8(y9, out7, out8)
	go acc9(y10, out8, out9)
	go acc10(y11, out9, out10)
	go acc11(y12, out10, out11)
	go acc12(y13, out11, out12)
	go acc13(y14, out12, out13)
	go acc14(y15, out13, out14)
	go acc15(y16, out14, out15)
	go acc16(y17, out15, out16)
	go acc17(y18, out16, out17)
	go out_control(y19, out17, y)

	go input_first(x, x0, in0)
	go input0(in0, x1, in1)
	go input1(in1, x2, in2)
	go input2(in2, x3, in3)
	go input3(in3, x4, in4)
	go input4(in4, x5, in5)
	go input5(in5, x6, in6)
	go input6(in6, x7, in7)
	go input7(in7, x8, in8)
	go input8(in8, x9, in9)
	go input9(in9, x10, in10)
	go input10(in10, x11, in11)
	go input11(in11, x12, in12)
	go input12(in12, x13, in13)
	go input13(in13, x14, in14)
	go input14(in14, x15, in15)
	go input15(in15, x16, in16)
	go input16(in16, x17, in17)
	go input17(in17, x18, in18)
	go input_last(in18, x19)
}