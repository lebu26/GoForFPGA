package main
 
import (
	//"image"
	//"image/color"
	//"image/png"
	"log"
	"os"
	"strconv"
	"time"
	//"github.com/ReconfigureIO/fixed"
	//"github.com/ReconfigureIO/fixed/host"
	"github.com/ReconfigureIO/sdaccel/xcl"
	"encoding/binary"
)

type Int16_16 int32

func (a Int16_16)Mul(b Int16_16) Int16_16{ //Int16_16
	return Int16_16((int64(a)*int64(b) + 1<<15) >> 16)
} 

func to_fix(a float64) Int16_16{
	return Int16_16(a*(1<<16))
}

func main() {
	arg_w := os.Args[1]
	width, error_w := strconv.Atoi(arg_w)

	if error_w == nil{

	}
	arg_h := os.Args[2]
	height, error_h := strconv.Atoi(arg_h)

	if error_h == nil{

	}

	log.Printf("**********************program starting*********************\n\r")
	log.Printf("Software part staring...\n\r")
	var t1_sw = time.Now()
	var width_i, height_i int32 = int32(width), int32(height)
	//width := fixed.Int52_12(width_i*64)
	//height := fixed.Int52_12(height_i*64)
	//var maxIter uint8 = 255
	var cX_f, cY_f float64 = -0.7, 0.27
	//fileName := "julia_fixed.png"
	//img := image.NewNRGBA(image.Rect(0, 0, int(width_i), int(height_i)))
	cX := to_fix(cX_f)
 	cY := to_fix(cY_f)
 	para_w:=to_fix(2.0/float64(width_i))
 	para_h:=to_fix(2.0/float64(height_i))
	gold  := make([]uint8, int(width_i*height_i))
	para  := make([]uint32, 4) //for parameters that requires divide
	para[0] = uint32(para_w)
	para[1] = uint32(para_h)
	para[2] = uint32(cX)
	para[3] = uint32(cY)
 
	var x int32
	f_1 := Int16_16(65536) //1
	f_1_5 := Int16_16(98304) //1.5
	f_2 := Int16_16(131072)//2
	f_4 := Int16_16(262144)//4

	var maxIter uint8 = 255

	//var x int32
	var y int32
	var thisx Int16_16
	var thisy Int16_16
	for x = 0; x < width_i; x++ {
		thisx = Int16_16(x*65536)
		var tmp, zx, zy Int16_16
		var i uint8
		for y = 0; y < height_i; y++ {
			thisy = Int16_16(y*65536)
			zx = f_1_5.Mul(thisx.Mul(para_w) - f_1)
			zy = thisy.Mul(para_h) - f_1
			i = maxIter
			
			zx2 := zx.Mul(zx)
			zy2 := zy.Mul(zy)
			
			for (zx2+zy2) < f_4 && i > 0 {
				//tmp = zx.Mul(zx) - zy.Mul(zy) - Int16_16(2867)
				tmp = zx2 - zy2 + cX
				zy = f_2.Mul(zx.Mul(zy)) + cY
				zx = tmp
				i--
				zx2 = zx.Mul(zx)
				zy2 = zy.Mul(zy)
			}
			//img.Set(int(x), int(y), color.RGBA{i, i, i << 3, 255})
			gold[x*height_i+y] = i
		}
	}
	/*
	imgFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer imgFile.Close()
	if err := png.Encode(imgFile, img); err != nil {
		imgFile.Close()
		log.Fatal(err)
	}*/
	var t2_sw = time.Now()
	td_sw := t2_sw.Sub(t1_sw)
	log.Printf("**********************software finished*********************\n\r\r\r")
	log.Printf("**software takes = %v s\n\r", td_sw.Seconds())
	log.Printf("**cx = %v \n\r", int32(para[2]))
	log.Printf("**cy = %v \n\r", para[3])
	log.Printf("**para_w = %v \n\r",para[0])
	log.Printf("**para_h = %v \n\r", para[1])

	log.Printf("************************************************************************************* \n\r")
	log.Printf("hardware part staring...\n\r")


	// Allocate a 'world' for interacting with the FPGA
	world := xcl.NewWorld()
	defer world.Release()

	// Import the compiled code that will be loaded onto the FPGA (referred to here as a kernel)
	// Right now these two identifiers are hard coded as an output from the build process
	krnl := world.Import("kernel_test").GetKernel("reconfigure_io_sdaccel_builder_stub_0_1")
	defer krnl.Release()

	// Construct an array to hold the output data from the FPGA
	
	buff_para := world.Malloc(xcl.ReadWrite, uint(binary.Size(para)))
	defer buff_para.Free()

	output := make([]uint32, width_i*height_i)
	// Allocate a space in the shared memory to store the output data from the FPGA.
	buff_out := world.Malloc(xcl.ReadWrite, uint(binary.Size(output)))
	defer buff_out.Free()

	debug := make([]uint32, 4)
	// Allocate a space in the shared memory to store the output data from the FPGA.
	buff_debug := world.Malloc(xcl.ReadWrite, uint(binary.Size(debug)))
	defer buff_debug.Free()

	// Zero out the space in shared memory for the result from the FPGA
	binary.Write(buff_out.Writer(), binary.LittleEndian, output)
	binary.Write(buff_para.Writer(), binary.LittleEndian, para)
	binary.Write(buff_debug.Writer(), binary.LittleEndian, debug)

	// Pass the pointer to the input data in shared memory as the first and second argument
	krnl.SetArg(0, uint32(width_i))
	krnl.SetArg(1, uint32(height_i))
	//Pass the pointer to the output data in shared memory as the third argument
	krnl.SetMemoryArg(2, buff_para)
	krnl.SetMemoryArg(3, buff_out)
	krnl.SetMemoryArg(4, buff_debug)

	var td_hw time.Duration
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
	err := binary.Read(buff_out.Reader(), binary.LittleEndian, output)
	if err != nil {
		log.Fatal("binary.Read failed:", err)
	}
	var t2_hw = time.Now()
	td_hw = t2_hw.Sub(t1_hw)

	err_count := 0
	var w int32
	for w=0 ; w<width_i*height_i; w++{
		index := int32(output[w]>>9)
		label := uint8(output[w]&(511))
		if label != gold[index]{
			err_count ++
			if err_count < 10{
				log.Printf("** output[%d] = %v, gold[%d] = %v \n\r", index,label,index,gold[index])
			}
		}	
	} 

	err_1 := binary.Read(buff_debug.Reader(), binary.LittleEndian, debug)
	if err_1 != nil {
		log.Fatal("binary.Read failed:", err_1)
	}

	//for w=0 ; w<width_i*height_i; w++{
	log.Printf("zx = %v \n\r",int32(debug[0]))
	log.Printf("zy = %v \n\r",int32(debug[1])) 
	log.Printf("pw = %v \n\r",int32(debug[2])) 
	log.Printf("ph = %v \n\r",int32(debug[3]))  
	//} 

	log.Printf("**********************hardware finished*********************\n\r\r\r")
	log.Printf("** takes = %v s\n\r", td_hw.Seconds())
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel finished successfully\n\r")
	log.Printf("                                  error count = %v \n\r",err_count)
	log.Printf("************************************************************************************* \n\r")
}