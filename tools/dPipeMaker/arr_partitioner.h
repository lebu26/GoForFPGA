#pragma once
#include <vector>
#include <string>
#include <regex>
#include <fstream>
#include <sstream>
#include "partitioner.h"
using namespace std;

struct arr_info
{
	vector<string> arr_element;//result vector
	string arr_name;
	string arr_type;
	int arr_length;
};


class arr_partitioner
{
private:
	partitioner p;//result vector
	arr_info info;
	string type = "block";
	int factor;
	int vector_size=1;
	string file_name;
	vector<string> res;
public:
	arr_partitioner() {}

	arr_partitioner(string _file_name, string cmd) 
	{//cmd in partern of #pragma partition value=xx factor=xx type=xx vector_size=xx
		file_name = _file_name;
		readCmd(cmd);

		findArray();
		p = partitioner(info.arr_length, factor, vector_size);

		parti_b();
/*

		if (type == "cyclic")
			parti_c();
		else
			parti_b();*/
	}

	arr_partitioner(string _file_name, string _arr_name, int _factor,int _vector_size,string _type )
	{
		file_name = _file_name;
		info.arr_name = _arr_name;
		factor = _factor;
		vector_size = _vector_size;
		type = _type;
		//find array
		findArray();
		p = partitioner(info.arr_length, factor, vector_size);
		//do partition

		parti_b();
/*		
		if (type == "cyclic")
		parti_c();
		else
		parti_b();
		
*/		
	}

	void readCmd(string cmd)
	{
		regex fac("factor *= *([^ ]+) *"), val("value *= *([^ ]+) *"), 
			typ("type *= *([^ ]+) *"), vcs("vector_size *= *([^ ]+) *");
		//const char* c = cmd.c_str();
		//char t[100];
		smatch facM, valM,typM,vcsM;
		//find the name of the array
		//cmd is #pragma partition value=two_pow factor=2 type=cyclic
		//sscanf(c, "%*s %*s value=%s factor=%d", t, &temp.factor);
		bool foundFac = regex_search(cmd, facM, fac);
		bool foundVal = regex_search(cmd, valM, val);
		bool foundTyp = regex_search(cmd, typM, typ);
		bool foundVcs = regex_search(cmd, vcsM, vcs);

		if (foundVal)
			info.arr_name = valM[1].str();

		if (foundFac)
			factor = stoi(facM[1].str());

		if (foundTyp)
			type = typM[1].str();

		if (foundVcs)
			vector_size = stoi(vcsM[1].str());
	}

	void findArray()
	{
		ifstream myfile(file_name);
		string line;
		string declear_pat = info.arr_name + " *:= *[\[]([^\\]]+)[\\]]([^\{]+)[\{]([^\}]+)[\}]";
		//
		regex declear_reg(declear_pat);
		//size_t found;
		//find declearation
		string file_buff = "";
		while (getline(myfile, line))
		{
			file_buff += line;
		}
		smatch arr_info;
		string temp_ele;
		if (regex_search(file_buff, arr_info, declear_reg))
		{
			info.arr_length = stoi(arr_info[1].str());
			info.arr_type = arr_info[2].str();
			temp_ele = arr_info[3].str();
		}
		stringstream ss(temp_ele);
		string tp;
		while (getline(ss, tp, ','))
		{
			info.arr_element.push_back(tp);
		}

		myfile.close();
	}

	void parti_b()
	{
		int start = 0;
		for (int i = 0; i < factor; i++)
		{
			string tempRes = info.arr_name + " := [" + to_string(p.getRes(i)) + "]"+info.arr_type+"{";
			for (int j = 0; j < p.getRes(i); j++)
			{
				tempRes += info.arr_element[start + j]+", ";
			}
			tempRes += "}";
			res.push_back(tempRes);
			start += p.getRes(i);
		}
	}

	string getRes(int i)
	{
		return res[i];
	}

	vector<string> getRes()
	{
		return res;
	}
	//might add a few more setter and getter 
	arr_info getInfo()
	{
		return info;
	}
};