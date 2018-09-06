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

func julia(width uint32, height uint32, para_w Int16_16, para_h Int16_16, cX Int16_16, cY Int16_16, out chan<- uint8){
	f_1 := Int16_16(65536) //1
	f_1_5 := Int16_16(98304) //1.5
	f_2 := Int16_16(131072)//2
	f_4 := Int16_16(262144)//4

	var maxIter uint8 = 255

	var x uint32
	var y uint32
	var thisx Int16_16
	var thisy Int16_16
	for x = 0; x < width; x++ {
		thisx = Int16_16(x*65536)
		var tmp, zx, zy Int16_16
		var i uint8
		for y = 0; y < height; y++ {
			thisy = Int16_16(y*65536)
			tm := thisx.Mul(para_w)
			tm = tm - f_1
			zx = f_1_5.Mul(tm)
			zy = thisy.Mul(para_h) - f_1
			i = maxIter
			
			zx2 := zx.Mul(zx)
			zy2 := zy.Mul(zy)
			
			for (zx2+zy2) < f_4 && i > 0 {
				tmp = zx2 - zy2 + cX
				tm1 := zx.Mul(zy)
				zy = f_2.Mul(tm1) + cY
				//zy = f_2.Mul(zx.Mul(zy)) + cY
				zx = tmp
				i--
				zx2 = zx.Mul(zx)
				zy2 = zy.Mul(zy)
			}
			out <- i
		}
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

	out := make(chan uint8,2)
	paraChan := make(chan uint32,2)
	go aximemory.ReadBurstUInt32(memReadAddr, memReadData, true, paraAddr, 4, paraChan)

	para_w := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr))
	para_h := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr+uintptr(4)))
	cX := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr+uintptr(8)))
	cY := Int16_16(<-paraChan)//aximemory.ReadUInt32(memReadAddr, memReadData, true, paraAddr+uintptr(12)))

	go julia(width,height,para_w,para_h,cX,cY,out)

	aximemory.WriteUInt32(
				memWriteAddr, memWriteData, memWriteResp, true,
				debugAddr,uint32(cX))
	aximemory.WriteUInt32(
				memWriteAddr, memWriteData, memWriteResp, true,
				debugAddr+uintptr(4),uint32(cY))
	aximemory.WriteUInt32(
				memWriteAddr, memWriteData, memWriteResp, true,
				debugAddr+uintptr(8),uint32(para_w))
	aximemory.WriteUInt32(
				memWriteAddr, memWriteData, memWriteResp, true,
				debugAddr+uintptr(12),uint32(para_h))


	aximemory.WriteBurstUInt8(memWriteAddr, memWriteData, memWriteResp, true, outputAddr, width*height, out)
}