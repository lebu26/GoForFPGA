#pragma once
#include<string>
#include<vector>
#include<fstream>
using namespace std;
class FIR {
private:
	string data_t =  "int"; //default data type for the fir filter is int
	int order = 1; //default order is 1
	vector<int>coef;
public:
	FIR(int _order, vector<int>_coef ) {
		order = _order;
		coef = _coef;
	}
	FIR(vector<int>_coef) {
		order = _coef.size();
		coef = _coef;
	}
	FIR(string _type, int _order, vector<int>_coef) {
		order = _order;
		data_t = _type;
		coef = _coef;
	}

	void FIR_gen() {
		ofstream myfile;
		myfile.open("FIR_40.go");
		myfile.clear();
		
		myfile << "package FIR" << endl<<endl;
		//myfile << "import(" << endl;
		//myfile << "	\"_github.com/ReconfigureIO/sdaccel\" " << endl;
		//myfile << "	aximemory \"github.com/ReconfigureIO/sdaccel/axi/memory\" " << endl;
		//myfile << "	axiprotocol \"github.com/ReconfigureIO/sdaccel/axi/protocol\"" << endl;
		//myfile << ")" << endl << endl;
/*
		myfile << "func acc(y0 <-chan " << data_t << "," << endl;
		for (int i = 1; i < order; i++) {
			myfile << "	y" << i << " <-chan " << data_t << "," << endl;
		}
		myfile << "	y chan<- uint32){" << endl;
		myfile << "	for{" << endl;
		for (int i = 0; i < order; i++) {
			myfile << "		res" << i << " := <-y" << i << endl;
		}
		myfile << "		y <- uint32(";
		for (int i = 0; i < order; i++) {
			myfile << "res" << i;
			if (i != order - 1) {
				myfile << " + ";
			}
		}
		myfile << ")" << endl << "	}" << endl << "}" << endl;
*/
		for (int i = 0; i < order - 2; i++) {
			myfile << "func acc"<< i <<"(y1 <-chan int, y2 <-chan int, out chan<- int){" << endl;
			myfile << "	for{" << endl;
			myfile << "		op1 := <-y1" << endl;
			myfile << "		op2 := <-y2" << endl;
			myfile << "		out <- op1 + op2" << endl;
			myfile << "	}" << endl << "}" << endl;
		}
		myfile << "func out_control(y1 <-chan int, y2 <-chan int, out chan<- uint32){" << endl;
		myfile << "	for{" << endl;
		myfile << "		op1 := <-y1" << endl;
		myfile << "		op2 := <-y2" << endl;
		myfile << "		out <- uint32(op1 + op2)" << endl;
		myfile << "	}" << endl << "}" << endl;

		for (int i = 0; i < order; i++) {
			myfile << "func fir" << i<<"(x"<<i<<" <-chan "<<data_t<<", y"<<i<<" chan<- "<<data_t<<"){"<<endl;
			myfile << "	for{" << endl;
			myfile << "		temp := " << data_t << "(<-x" << i << ")" << endl;
			myfile << "		y" << i << " <- temp*" << coef[i] << endl;
			myfile << "	}"<<endl;
			myfile << "}"<<endl;
		}
/*
		myfile << "func input_control(x0 chan<- " << data_t << "," << endl;
		for (int i = 1; i < order; i++) {
			myfile << "	x" << i << " chan<- " << data_t << "," << endl;
		}
		myfile << "	x <-chan uint32){" << endl;
		for (int i = 1; i < order; i++) {
			myfile << "		X" << i << " := 0" << endl;
		}
		myfile << "	for{" << endl;
		myfile << "		temp := " << data_t << "(<-x)" << endl;
		myfile << "		x0 " << "<- temp" << endl;
		for (int i = 1; i < order; i++) {
			myfile << "		x" << i << " <- X" << i << endl;
		}
		myfile << endl;
		for (int i = order-1; i > 1; i--) {
			myfile << "		X" << i << " = X"<<i-1 << endl;
		}
		myfile << "		X1 = temp" << endl;
		myfile << "	}" << endl << "}" << endl;
*/
		myfile << "func input_first(in <-chan uint32, out chan<- int, move chan<- int){" << endl;
		myfile << "	var x int = 0" << endl;
		myfile << "	for{" << endl;
		myfile << "		x = int(<-in)" << endl;
		myfile << "		out <-x" << endl;
		myfile << "		move <- x" << endl;
		myfile << "	}" << endl << "}" << endl;

		for (int i = 0; i < order - 2; i++) {
			myfile << "func input"<< i <<"(in <-chan int, out chan<- int, move chan<- int){" << endl;
			myfile << "	var x int = 0" << endl;
			myfile << "	for{" << endl;
			myfile << "		out <-x" << endl;
			myfile << "		move <- x" << endl;
			myfile << "		x = <-in" << endl;
			myfile << "	}" << endl << "}" << endl;
		}
		myfile << "func input_last(in <-chan int, out chan<- int){" << endl;
		myfile << "	var x int = 0" << endl;
		myfile << "	for{" << endl;
		myfile << "		out <-x" << endl;
		myfile << "		x = <-in" << endl;
		myfile << "	}" << endl << "}" << endl;

		myfile << "func FIR_Top(x <-chan uint32, y chan<- uint32){" << endl;

		for (int i = 0; i < order; i++) {
			myfile << "	x" << i << " := make(chan " << data_t << ", 2)" << endl;
			myfile << "	y" << i << " := make(chan " << data_t << ", 2)" << endl<<endl;
		}

		for (int i = 0; i < order-2; i++) {
			myfile << "	out" << i << " := make(chan " << data_t << ", 2)" << endl;
		}
		myfile << endl;

		for (int i = 0; i < order - 1; i++) {
			myfile << "	in" << i << " := make(chan " << data_t << ", 2)" << endl;
		}
		myfile << endl;

		for (int i = 0; i < order; i++) {
			myfile << "	go fir" << i << "(x" << i << ", y" << i << ")" << endl;
		}
		myfile << endl;

		myfile << "	go acc0(y0, y1, out0)" << endl;
		for (int i = 2; i < order-1; i++) {
			myfile << "	go acc" << i-1 << "(y" << i << ", out" << i-2 << ", out" << i - 1 << ")" << endl;
		}
		myfile << "	go out_control(y" << order-1 << ", out" << order - 3 <<", y)" << endl;
		myfile << endl;

		myfile << "	go input_first(x, x0, in0)" << endl;
		for (int i = 0; i < order - 2; i++) {
			myfile << "	go input" << i << "(in" << i << ", x" << i+1 << ", in" << i+1 << ")" << endl;
		}
		myfile << "	go input_last" << "(in" << order-2 << ", x" << order-1 << ")" << endl;
/*
		myfile << "	go acc(";
		for (int i = 0; i < order; i++) {
			myfile << "y" << i << ", ";
		}
		myfile << "y)" << endl;

		myfile << "	go input_control(";
		for (int i = 0; i < order; i++) {
			myfile << "x" << i << ", ";
		}
		myfile << "x)" << endl;
/*
		myfile << "	go func(){" << endl;
		myfile << "		for{" << endl;
		for (int i = 0; i < order; i++) {
			myfile << "			res" << i << " := <-y" << i << endl;
		}
		myfile << "			y <- uint32(";
		for (int i = 0; i < order; i++) {
			myfile << "res" << i;
			if (i != order - 1) {
				myfile << " + ";
			}
		}
		myfile << ")" << endl << "	}" << endl << "}()" << endl;

		myfile << "	go func(){" << endl;
		for (int i = 1; i < order; i++) {
			myfile << "		X" << i << " := 0" << endl;
		}
		myfile << "	for{" << endl;
		myfile << "		temp := " << data_t << "(<-x)" << endl;
		myfile << "		x0 " << "<- temp" << endl;
		for (int i = 1; i < order; i++) {
			myfile << "		x" << i << " <- X" << i << endl;
		}
		myfile << endl;
		for (int i = order - 1; i > 1; i--) {
			myfile << "		X" << i << " = X" << i - 1 << endl;
		}
		myfile << "		X1 = temp" << endl;
		myfile << "	}" << endl << "}()" << endl;
*/
		myfile << "}";

		myfile.close();
	}

};