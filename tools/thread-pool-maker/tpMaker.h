#pragma once
#include<string>
#include<vector>
#include<fstream>
using namespace std;
struct para {
	string name;
	string type;
};
/*The template is format as following
//functions for threads
func <funcName>(<inputChan>Chan <-chan type, <parameters>, <outChan>Chan chan<- type){
	for{
		<inputChan> := <- <inputChans>Chan
		...

		<func_body>
		...

		<outChan>Chan <- <outputChan> //please be noted that needs to encode the result with with labels
	}
}

//channel used declearation
declear of <inputChan>
declear of <outChan>
declear of final_out_chan

//go routines
go <funcName>(<inputChans>, <parameters>, <outChan>)
...

go routine for feeding input chan
go func(){
	var tm<inputChans>
	i := 0
	for{
		switch i{
			case i: 
				<inputChans> <- tm<inputChans>
				i++
			...
			default:
				i=0
		}
	}
}()

//function for collecting output
go func(){
	var res uint32
	for{
		select{
			case res = uint32(<- <outchan>) :
			...
		}
		final_out_chan <- res
	}
}()

*/
using namespace std;
class tpMaker {
private:
	int numOfTp = 1;
	string funcName;
	vector<para> inputChan;
	vector<para> parameters;
	vector<para> outChan;
	string funcBody;
public: 
	tpMaker() {}

	void setNumOfTp(int _numOfTop) {
		numOfTp = _numOfTop;
	}

	void setInputChan(string _name, string _type) {
		para p = { _name, _type };
		inputChan.push_back(p);
	}
	void setPara(string _name, string _type) {
		para p = { _name, _type };
		parameters.push_back(p);
	}
	void setOutputChan(string _name, string _type) {
		para p = { _name, _type };
		outChan.push_back(p);
	}
	void setFuncName(string _name) {
		funcName = _name;
	}
	void setFuncbody(string _body) {
		funcBody = _body;
	}

	void gen() {
		ofstream file;
		string fileName = "RecoTP_" + funcName + ".go";
		file.open(fileName);
		file << "package TreadPool" << endl;
		//generate TP functions
		file << "//***********TP functions************" << endl;
		for (int i = 0; i < numOfTp; i++) {
			//function head
			file << "func " << funcName <<i<<"(";
			for (int j = 0; j < inputChan.size(); j++) {
				file << inputChan[j].name << "Chan <-chan " << inputChan[j].type << ", ";
			}
			for (int j = 0; j < parameters.size(); j++) {
				file << parameters[j].name << " " << parameters[j].type << ", ";
			}
			for (int j = 0; j < outChan.size()-1; j++) {
				file << outChan[j].name << "Chan chan<- " << outChan[j].type << ", ";
			}
			if (outChan.size() > 0) {
				file << outChan[outChan.size()-1].name << "Chan chan<- " << outChan[outChan.size()-1].type << "){ "<<endl;
			}

			//function body
			file << "	for{" << endl;
			for (int j = 0; j < inputChan.size(); j++) {
				file << "		"<< inputChan[j].name << " := <-" << inputChan[j].name <<"Chan "<<endl;
			}
			file << "//function body" << endl;
			file << funcBody <<endl;
			
			//handle output
			for (int j = 0; j < outChan.size(); j++) {
				file << "		" << outChan[j].name << "Chan <-" << outChan[j].name << endl;
			}

			file << "	}" << endl << "}" << endl;
		}//end of TP functions
		
		//generate the top level function
		file << endl << "//***********Top level function************" << endl;
		file << "func Top(";
		for (int j = 0; j < inputChan.size(); j++) {
			file << inputChan[j].name << "Chan <-chan " << inputChan[j].type << ", ";
		}
		for (int j = 0; j < parameters.size(); j++) {
			file << parameters[j].name << " " << parameters[j].type << ", ";
		}
		for (int j = 0; j < outChan.size() - 1; j++) {
			file << outChan[j].name << "Chan chan<- " << outChan[j].type << ", ";
		}
		if (outChan.size() > 0) {
			file << outChan[outChan.size() - 1].name << "Chan chan<- " << outChan[outChan.size() - 1].type << "){ " << endl;
		}

		//generate channel declearation
		file <<endl << "//***********channels declearation************" << endl;
		file << endl;
		for (int i = 0; i < numOfTp; i++) {
			for (int j = 0; j < inputChan.size(); j++) {
				file << "	" << inputChan[j].name << "Chan"<< i <<" := make(chan " << inputChan[j].type << ", 2)" << endl;
			}
			for (int j = 0; j < outChan.size(); j++) {
				file << "	" << outChan[j].name << "Chan" << i << " := make(chan " << outChan[j].type << ", 2)" << endl;
			}
			file << endl;
		}//end of channel declear

		 //generate input go routines
		file << endl << "//***********input fan go routines declearation************" << endl;
		file << "	go func(){ //for feed data in" << endl;
		file << "		i:=0" << endl;
		file << "		for{" << endl;
		for (int j = 0; j < inputChan.size(); j++) {
			file << "		tmp_" << inputChan[j].name << " := <-" << inputChan[j].name << "Chan " << endl;
		}
		file << "			switch i{"<<endl;

		for (int i = 0; i < numOfTp-1; i++) {
			file << "				case "<<i<<":"<<endl;
			for (int j = 0; j < inputChan.size(); j++) {
				file << "				" << inputChan[j].name << "Chan" << i << " <- tmp_" << inputChan[j].name << endl;
			}
			file << "				i++" << endl;
		}
		file << "				default :" << endl;
		for (int j = 0; j < inputChan.size(); j++) {
			file << "				" << inputChan[j].name << "Chan" << numOfTp-1 << " <- tmp_" << inputChan[j].name << endl;
		}
		file << "				i=0" << endl;
		file << "			}" << endl;
		file << "		}" << endl;
		file << "	}()" << endl;
	

		//generate TP go routines
		file << endl << "//***********TP go routines declearation************" << endl;
		for (int i = 0; i < numOfTp; i++) {
			//function head
			file << "	go " << funcName << i << "(";
			for (int j = 0; j < inputChan.size(); j++) {
				file << inputChan[j].name << "Chan"<<i<<", ";
			}
			for (int j = 0; j < parameters.size(); j++) {
				file << parameters[j].name << ", ";
			}
			for (int j = 0; j < outChan.size() - 1; j++) {
				file << outChan[j].name << "Chan" << i << ", ";
			}
			if (outChan.size() > 0) {
				file << outChan[outChan.size() - 1].name << "Chan" << i << ")" << endl;
			}
		}

		//generate output go routines
		file << endl << "//***********collecting output data from threads************" << endl;
		for (int a = 0; a < outChan.size(); a++) {
			file << "	go func(){ //for collecting data and out put" << endl;
			file << "		var" <<outChan[a].name<<"_res "<<outChan[a].type<< endl;
			file << "		for{" << endl;
			file << "			select {"<<endl;

			for (int i = 0; i < numOfTp; i++) {
				file << "				case " << outChan[a].name << "_res = <-"<<outChan[a].name << "Chan"<<i<<":" << endl;
			}
			file << "			}" << endl;
			file << "			" << outChan[a].name << "Chan <- " << outChan[a].name << "_res" << endl;
			file << "		}" << endl;
			file << "	}()" << endl;
		}

		file << "}" << endl;//end of the top-level function
	}//end of gen()

};//end of class