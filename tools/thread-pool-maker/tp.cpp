#include<iostream>
#include"tpMaker.h"
using namespace std;

void main() {
	tpMaker tp;
	//set name of the function
	tp.setFuncName("julia");

	tp.setNumOfTp(10);

	//set input channels for top level
	tp.setInputChan("x", "uint32");
	tp.setInputChan("y", "uint32");

	//set parameters for top level
	tp.setPara("height", "uint32");
	tp.setPara("width", "uint32");
	tp.setPara("para_w", "Int16_16");
	tp.setPara("para_h", "Int16_16");
	tp.setPara("cX", "Int16_16");
	tp.setPara("cY", "Int16_16");

	//set function body
	string func_body = "label := x*height + y \n \
		thisx := Int16_16(x * 65536) \n \
		thisy := Int16_16(y * 65536)\n \
		tm := thisx.Mul(para_w)\n \
		tm = tm - 65536 \n \
		f_1_5 : = Int16_16(98304) \n \
		zx := f_1_5.Mul(tm)\n \
		zy := thisy.Mul(para_h) - 65536\n \
		i := uint32(255)\n \
\n \
		zx2 := zx.Mul(zx)\n \
		zy2 := zy.Mul(zy)\n \
		f_4 := Int16_16(262144)\n \
		f_2 := Int16_16(131072)\n \
\n \
		for (zx2 + zy2) < f_4 && i > 0 {\n \
		tmp: = zx2 - zy2 + cX\n \
			tm1 := zx.Mul(zy)\n \
			zy = f_2.Mul(tm1) + cY\n \
			//zy = f_2.Mul(zx.Mul(zy)) + cY\n \
			zx = tmp\n \
			i--\n \
			zx2 = zx.Mul(zx)\n \
			zy2 = zy.Mul(zy)\n \
		}\n \
	//encode result\n \
	out := uint32(label << 9 + i)";
	tp.setFuncbody(func_body);

	//set output variable
	tp.setOutputChan("out", "uint32");
	
	//generate the file
	tp.gen();
}