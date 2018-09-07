package main
import(
	"_github.com/ReconfigureIO/sdaccel" 
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory" 
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"
)

func acc(y0 <-chan int,
	y1 <-chan int,
	y2 <-chan int,
	y3 <-chan int,
	y4 <-chan int,
	y5 <-chan int,
	y6 <-chan int,
	y7 <-chan int,
	y8 <-chan int,
	y9 <-chan int,
	y chan<- uint32){
	for{
		res0 := <-y0
		res1 := <-y1
		res2 := <-y2
		res3 := <-y3
		res4 := <-y4
		res5 := <-y5
		res6 := <-y6
		res7 := <-y7
		res8 := <-y8
		res9 := <-y9
		y <- uint32(res0 + res1 + res2 + res3 + res4 + res5 + res6 + res7 + res8 + res9)
	}
}
func fir0(x0 <-chan int, y0 chan<- int){
	for{
		temp := int(<-x0)
		y0 <- temp*1
	}
}
func fir1(x1 <-chan int, y1 chan<- int){
	for{
		temp := int(<-x1)
		y1 <- temp*2
	}
}
func fir2(x2 <-chan int, y2 chan<- int){
	for{
		temp := int(<-x2)
		y2 <- temp*3
	}
}
func fir3(x3 <-chan int, y3 chan<- int){
	for{
		temp := int(<-x3)
		y3 <- temp*4
	}
}
func fir4(x4 <-chan int, y4 chan<- int){
	for{
		temp := int(<-x4)
		y4 <- temp*5
	}
}
func fir5(x5 <-chan int, y5 chan<- int){
	for{
		temp := int(<-x5)
		y5 <- temp*6
	}
}
func fir6(x6 <-chan int, y6 chan<- int){
	for{
		temp := int(<-x6)
		y6 <- temp*7
	}
}
func fir7(x7 <-chan int, y7 chan<- int){
	for{
		temp := int(<-x7)
		y7 <- temp*8
	}
}
func fir8(x8 <-chan int, y8 chan<- int){
	for{
		temp := int(<-x8)
		y8 <- temp*9
	}
}
func fir9(x9 <-chan int, y9 chan<- int){
	for{
		temp := int(<-x9)
		y9 <- temp*10
	}
}
func input_control(x0 chan<- int,
	x1 chan<- int,
	x2 chan<- int,
	x3 chan<- int,
	x4 chan<- int,
	x5 chan<- int,
	x6 chan<- int,
	x7 chan<- int,
	x8 chan<- int,
	x9 chan<- int,
	x <-chan uint32){
		X1 := 0
		X2 := 0
		X3 := 0
		X4 := 0
		X5 := 0
		X6 := 0
		X7 := 0
		X8 := 0
		X9 := 0
	for{
		temp := int(<-x)
		x0 <- temp
		x1 <- X1
		x2 <- X2
		x3 <- X3
		x4 <- X4
		x5 <- X5
		x6 <- X6
		x7 <- X7
		x8 <- X8
		x9 <- X9

		X9 = X8
		X8 = X7
		X7 = X6
		X6 = X5
		X5 = X4
		X4 = X3
		X3 = X2
		X2 = X1
		X1 = temp
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
	go acc(y0, y1, y2, y3, y4, y5, y6, y7, y8, y9, y)
	go input_control(x0, x1, x2, x3, x4, x5, x6, x7, x8, x9, x)
}