package main
 
import (
	//"image"
	//"image/color"
	//"image/png"
	"log"
	//"os"
	"time"
	"github.com/ReconfigureIO/fixed"
	"github.com/ReconfigureIO/fixed/host"
	"github.com/ReconfigureIO/sdaccel/xcl"
	"encoding/binary"
)
 
func main() {
	log.Printf("**********************program starting*********************\n\r")
	log.Printf("Software part staring...\n\r")
	var t1_sw = time.Now()
	var width_i, height_i int32 = 2, 2
	//width := fixed.Int52_12(width_i*64)
	//height := fixed.Int52_12(height_i*64)
	var maxIter uint8 = 255
	var cX_f, cY_f float64 = -0.7, 0.27
	//fileName := "julia_fixed.png"
	//img := image.NewNRGBA(image.Rect(0, 0, int(width_i), int(height_i)))
	cX := host.I52Float64(cX_f)
 	cY := host.I52Float64(cY_f)
	gold  := make([]uint8, int(width_i*height_i))
	para  := make([]uint64, 4) //for parameters that requires divide
	para[0] = uint64(host.I52Float64(2.0/float64(width_i)))
	para[1] = uint64(host.I52Float64(2.0/float64(height_i)))
	para[2] = uint64(cX)
	para[3] = uint64(cY)
 
	var x int32
	for x = 0; x < width_i; x++ {
		thisx := host.I52Float64(float64(x))
		var tmp, zx, zy fixed.Int52_12
		var i uint8
		var y int32
		for y = 0; y < height_i; y++ {
			thisy := host.I52Float64(float64(y))
			zx = host.I52Float64(float64(1.5)).Mul((thisx.Mul(host.I52Float64(2.0/float64(width_i))) - host.I52Float64(float64(1))))
			zy = thisy.Mul(host.I52Float64(2.0/float64(height_i))) - host.I52Float64(float64(1))
			i = maxIter
			for zx.Mul(zx)+zy.Mul(zy) < host.I52Float64(float64(4)) && i > 0 {
				tmp = zx.Mul(zx) - zy.Mul(zy) + cX
				zy = host.I52Float64(float64(2)).Mul(zx.Mul(zy)) + cY
				zx = tmp
				i--
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
	log.Printf("**cx = %v \n\r", para[2])
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

	output := make([]uint8, width_i*height_i)
	// Allocate a space in the shared memory to store the output data from the FPGA.
	buff_out := world.Malloc(xcl.ReadWrite, uint(binary.Size(output)))
	defer buff_out.Free()
/*
	debug := make([]uint64, width_i*height_i)
	// Allocate a space in the shared memory to store the output data from the FPGA.
	buff_debug := world.Malloc(xcl.ReadWrite, uint(binary.Size(debug)))
	defer buff_debug.Free()
*/
	// Zero out the space in shared memory for the result from the FPGA
	binary.Write(buff_out.Writer(), binary.LittleEndian, output)
	binary.Write(buff_para.Writer(), binary.LittleEndian, para)
//	binary.Write(buff_debug.Writer(), binary.LittleEndian, debug)

	// Pass the pointer to the input data in shared memory as the first and second argument
	krnl.SetArg(0, uint32(width_i))
	krnl.SetArg(1, uint32(height_i))
	//Pass the pointer to the output data in shared memory as the third argument
	krnl.SetMemoryArg(2, buff_para)
	krnl.SetMemoryArg(3, buff_out)
//	krnl.SetMemoryArg(4, buff_debug)

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
		if output[w] != gold[w]{
			err_count ++
			if err_count < 10{
				log.Printf("** output[%d] = %v, gold[%d] = %v \n\r", w,int32(output[w]),w,gold[w])
			}
		}	
	} 
	/*
	err_1 := binary.Read(buff_debug.Reader(), binary.LittleEndian, debug)
	if err_1 != nil {
		log.Fatal("binary.Read failed:", err_1)
	}

	for w=0 ; w<width_i*height_i; w++{
		log.Printf("%d,  zy = %v \n\r", w,fixed.Int52_12(debug[w])) 
	} 
	*/
	log.Printf("**********************hardware finished*********************\n\r\r\r")
	log.Printf("** takes = %v s\n\r", td_hw.Seconds())
	log.Printf("************************************************************************************* \n\r")
	log.Printf("                                  kernel finished successfully\n\r")
	log.Printf("                                  error count = %v \n\r",err_count)
	log.Printf("************************************************************************************* \n\r")

}