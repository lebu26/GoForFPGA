package main

import (
	// Import the entire framework for interracting with SDAccel from Go (including bundled verilog)
	_ "github.com/ReconfigureIO/sdaccel"
	
  // Use the new AXI protocol package for interracting with memory
	aximemory "github.com/ReconfigureIO/sdaccel/axi/memory"
	axiprotocol "github.com/ReconfigureIO/sdaccel/axi/protocol"

	//"github.com/ReconfigureIO/fixed"
)
type Int16_16 int

func (a Int16_16)Mul(b Int16_16) Int16_16{ //Int16_16
	return Int16_16((int64(a)*int64(b) + 1<<15) >> 16)
} 

func julia0(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia1(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia2(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia3(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia4(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia5(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia6(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia7(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)
		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia8(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia9(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia10(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia11(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia12(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia13(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia14(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia15(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia16(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia17(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)
		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia18(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}

func julia19(xchan <-chan uint32, ychan <-chan uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, outchan chan<- uint32){
	for{
		x := <-xchan
		y := <-ychan
		label := x*height + y
		thisx := Int16_16(x*65536)
		thisy := Int16_16(y*65536)
		tm := thisx.Mul(para_w)
		tm = tm - 65536
		f_1_5:=Int16_16(98304)
		zx := f_1_5.Mul(tm)
		zy := thisy.Mul(para_h) - 65536
		i := uint32(255)
		
		zx2 := zx.Mul(zx)
		zy2 := zy.Mul(zy)
		f_4 := Int16_16(262144)
		f_2 := Int16_16(131072)

		for (zx2+zy2) < f_4 && i > 0 {
			tmp := zx2 - zy2 + cX
			tm1 := zx.Mul(zy)
			zy = f_2.Mul(tm1) + cY
			//zy = f_2.Mul(zx.Mul(zy)) + cY
			zx = tmp
			i--
			zx2 = zx.Mul(zx)
			zy2 = zy.Mul(zy)
		}
		//encode result
		outchan <- (label<<9 + i)
	}
}
func Top(
	// Three operands from the host. Pointers to the input data and the space for the result in shared
	// memory and the length of the input data so the FPGA knows what to expect.
	width uint32,
	height uint32,
	paraAddr uintptr,
	outputAddr uintptr,
	debugAddr uintptr,

	// Set up channels for interacting with the shared memory
	memReadAddr chan<- axiprotocol.Addr,
	memReadData <-chan axiprotocol.ReadData,

	memWriteAddr chan<- axiprotocol.Addr,
	memWriteData chan<- axiprotocol.WriteData,
	memWriteResp <-chan axiprotocol.WriteResp) {

	out := make(chan uint32,2)
	paraChan := make(chan uint32,2)
	zxChan := make(chan uint32,2)
	zyChan := make(chan uint32,2)

	xchan0 := make(chan uint32,2)
	ychan0 := make(chan uint32,2)
	xchan1 := make(chan uint32,2)
	ychan1 := make(chan uint32,2)
	xchan2 := make(chan uint32,2)
	ychan2 := make(chan uint32,2)
	xchan3 := make(chan uint32,2)
	ychan3 := make(chan uint32,2)
	xchan4 := make(chan uint32,2)
	ychan4 := make(chan uint32,2)
	xchan5 := make(chan uint32,2)
	ychan5 := make(chan uint32,2)
	xchan6 := make(chan uint32,2)
	ychan6 := make(chan uint32,2)
	xchan7 := make(chan uint32,2)
	ychan7 := make(chan uint32,2)
	xchan8 := make(chan uint32,2)
	ychan8 := make(chan uint32,2)
	xchan9 := make(chan uint32,2)
	ychan9 := make(chan uint32,2)
	xchan10 := make(chan uint32,2)
	ychan10 := make(chan uint32,2)
	xchan11 := make(chan uint32,2)
	ychan11 := make(chan uint32,2)
	xchan12 := make(chan uint32,2)
	ychan12 := make(chan uint32,2)
	xchan13 := make(chan uint32,2)
	ychan13 := make(chan uint32,2)
	xchan14 := make(chan uint32,2)
	ychan14 := make(chan uint32,2)
	xchan15 := make(chan uint32,2)
	ychan15 := make(chan uint32,2)
	xchan16 := make(chan uint32,2)
	ychan16 := make(chan uint32,2)
	xchan17 := make(chan uint32,2)
	ychan17 := make(chan uint32,2)
	xchan18 := make(chan uint32,2)
	ychan18 := make(chan uint32,2)
	xchan19 := make(chan uint32,2)
	ychan19 := make(chan uint32,2)

	outchan0 := make(chan uint32,2)
	outchan1 := make(chan uint32,2)
	outchan2 := make(chan uint32,2)
	outchan3 := make(chan uint32,2)
	outchan4 := make(chan uint32,2)
	outchan5 := make(chan uint32,2)
	outchan6 := make(chan uint32,2)
	outchan7 := make(chan uint32,2)
	outchan8 := make(chan uint32,2)
	outchan9 := make(chan uint32,2)
	outchan10 := make(chan uint32,2)
	outchan11 := make(chan uint32,2)
	outchan12 := make(chan uint32,2)
	outchan13 := make(chan uint32,2)
	outchan14 := make(chan uint32,2)
	outchan15 := make(chan uint32,2)
	outchan16 := make(chan uint32,2)
	outchan17 := make(chan uint32,2)
	outchan18 := make(chan uint32,2)
	outchan19 := make(chan uint32,2)

	go aximemory.ReadBurstUInt32(memReadAddr, memReadData, true, paraAddr, 4, paraChan)

	para_w := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr))
	para_h := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr+uintptr(4)))
	cX := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr+uintptr(8)))
	cY := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr+uintptr(12)))

	go func(){
		for i:=uint32(0); i<width ; i++{
			for j:=uint32(0); j<height;j++{
				zxChan <- i
				zyChan <- j
			}
		}
	}()

	go julia0(xchan0, ychan0, height, para_w, para_h, cX, cY, outchan0)
	go julia1(xchan1, ychan1, height, para_w, para_h, cX, cY, outchan1)
	go julia2(xchan2, ychan2, height, para_w, para_h, cX, cY, outchan2)
	go julia3(xchan3, ychan3, height, para_w, para_h, cX, cY, outchan3)
	go julia4(xchan4, ychan4, height, para_w, para_h, cX, cY, outchan4)
	go julia5(xchan5, ychan5, height, para_w, para_h, cX, cY, outchan5)
	go julia6(xchan6, ychan6, height, para_w, para_h, cX, cY, outchan6)
	go julia7(xchan7, ychan7, height, para_w, para_h, cX, cY, outchan7)
	go julia8(xchan8, ychan8, height, para_w, para_h, cX, cY, outchan8)
	go julia9(xchan9, ychan9, height, para_w, para_h, cX, cY, outchan9)
	go julia10(xchan10, ychan10, height, para_w, para_h, cX, cY, outchan10)
	go julia11(xchan11, ychan11, height, para_w, para_h, cX, cY, outchan11)
	go julia12(xchan12, ychan12, height, para_w, para_h, cX, cY, outchan12)
	go julia13(xchan13, ychan13, height, para_w, para_h, cX, cY, outchan13)
	go julia14(xchan14, ychan14, height, para_w, para_h, cX, cY, outchan14)
	go julia15(xchan15, ychan15, height, para_w, para_h, cX, cY, outchan15)
	go julia16(xchan16, ychan16, height, para_w, para_h, cX, cY, outchan16)
	go julia17(xchan17, ychan17, height, para_w, para_h, cX, cY, outchan17)
	go julia18(xchan18, ychan18, height, para_w, para_h, cX, cY, outchan18)
	go julia19(xchan19, ychan19, height, para_w, para_h, cX, cY, outchan19)

	go func (){
		i := 0
		for{
			zx := <-zxChan
			zy := <-zyChan
			switch i{
			case 0:
				xchan0 <- zx
				ychan0 <- zy
				i++
			case 1:
				xchan1 <- zx
				ychan1 <- zy
				i++
			case 2:
				xchan2 <- zx
				ychan2 <- zy
				i++
			case 3:
				xchan3 <- zx
				ychan3 <- zy
				i++
			case 4:
				xchan4 <- zx
				ychan4 <- zy
				i++
			case 5:
				xchan5 <- zx
				ychan5 <- zy
				i++
			case 6:
				xchan6 <- zx
				ychan6 <- zy
				i++
			case 7:
				xchan7 <- zx
				ychan7 <- zy
				i++
			case 8:
				xchan8 <- zx
				ychan8 <- zy
				i++
			case 9:
				xchan9 <- zx
				ychan9 <- zy
				i++
			case 10:
				xchan10 <- zx
				ychan10 <- zy
				i++
			case 11:
				xchan11 <- zx
				ychan11 <- zy
				i++
			case 12:
				xchan12 <- zx
				ychan12 <- zy
				i++
			case 13:
				xchan13 <- zx
				ychan13 <- zy
				i++
			case 14:
				xchan14 <- zx
				ychan14 <- zy
				i++
			case 15:
				xchan15 <- zx
				ychan15 <- zy
				i++
			case 16:
				xchan16 <- zx
				ychan16 <- zy
				i++
			case 17:
				xchan17 <- zx
				ychan17 <- zy
				i++
			case 18:
				xchan18 <- zx
				ychan18 <- zy
				i++
			default:
				xchan19 <- zx
				ychan19 <- zy
				i = 0
			}
		}
	}()

	go func(){
		var res uint32
		for{
			select{
				case res = <-outchan0:
				case res = <-outchan1:
				case res = <-outchan2:
				case res = <-outchan3:
				case res = <-outchan4:
				case res = <-outchan5:
				case res = <-outchan6:
				case res = <-outchan7:
				case res = <-outchan8:
				case res = <-outchan9:
				case res = <-outchan10:
				case res = <-outchan11:
				case res = <-outchan12:
				case res = <-outchan13:
				case res = <-outchan14:
				case res = <-outchan15:
				case res = <-outchan16:
				case res = <-outchan17:
				case res = <-outchan18:
				case res = <-outchan19:
			}
			out <- res
		}
	}()

	aximemory.WriteBurstUInt32(memWriteAddr, memWriteData, memWriteResp, true, outputAddr, width*height, out)
}