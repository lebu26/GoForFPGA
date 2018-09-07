#pragma once
#include <vector>
using namespace std;
class partitioner
{
	private:
		int total;//total number
		int factor;//partition factor
		vector<int> res;//result vector
		int vector_size = 1; //default vector size is 1
	public:
		partitioner() {}
		partitioner(int t, int f)
		{
			total = t;
			factor = f;
			int subsize = t / f;
			int residue = t % f;
			for (int i = 0; i < f; i++)
			{
				int cur_size = residue > 0 ? subsize + 1 : subsize;
				res.push_back(cur_size);
				residue--;
			}
		}
		partitioner(int t, int f, int vs)
		{
			total = t;
			factor = f;
			vector_size = vs;
			int subsize = (total/vector_size) / factor;
			int residue = (total/vector_size) % factor;
			for (int i = 0; i < factor; i++)
			{
				int cur_size = residue > 0 ? subsize + 1 : subsize;
				cur_size *= vector_size;
				res.push_back(cur_size);
				residue--;
			}
		}


		int getRes(int i)
		{
			return res[i];
		}

		vector<int> getRes()
		{
			return res;
		}
		//might add a few more setter and getter 

};