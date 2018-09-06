
// kernel:  vadd
#define N 10
kernel __attribute__((reqd_work_group_size(1, 1, 1)))
void krnl_vadd(
                global const int* input,
                global int* output,
                const int length)
{
    int coef[N] = {1,2,3,4,5,6,7,8,9,10};
    int X[N] = {0,0,0,0,0,0,0,0,0,0};
    int order = N;
    int temp = 0;
    int i=0, j=0, w=0;

    __attribute__((xcl_pipeline_loop))
    loop1:for (i=0 ;i<length;i++){
        int temp =0;
        X[N-1] = input[i];
        __attribute__((xcl_pipeline_loop))
        loop2:for(int j=1;j<=N;j++){
            temp += X[N-j]*coef[j-1];
        }
        output[i] = temp;
        __attribute__((xcl_pipeline_loop))
        loop3:for(w=0;w<N-1;w++){
           X[w] = X[w+1];
        }
        /*
        X[0] = X[1];
        X[1] = X[2];
        X[2] = X[3];
        X[3] = X[4];
        X[4] = X[5];
        X[5] = X[6];
        X[6] = X[7];
        X[7] = X[8];
        X[8] = X[9];*/
    }
}
