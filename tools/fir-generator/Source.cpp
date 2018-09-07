#include"fir_gen.h"
#include <stdlib.h>
#include<fstream>

using namespace std;

int main() {
	int N = 40;
	ofstream file;
	file.open("coef_40.txt");

	vector<int>coef;
	for (int i = 0; i < N ; i++) {
		int tmp = rand() % 10 + 1;
		coef.push_back(tmp);
		file << tmp << ", ";
	}
	file.close();
	//int tmp = rand() % 10 + 1;
	//int myints[] = { 1,2,3,4,5,6,7,8,9,10 };
	//vector<int>coef(myints, myints + sizeof(myints) / sizeof(int));
	FIR f(coef);
	f.FIR_gen();
}