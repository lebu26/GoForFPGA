int mul(int a, int b){
    long la = (long)a;
    long lb = (long)b;
    long res = la*lb;
    return (int)((res+(1<<15))>>16);
}

kernel __attribute__((reqd_work_group_size(1, 1, 1)))
void krnl_vadd(
                global int* c,
                const int width,
                const int height,
                const int para_w,
                const int para_h,
                const int cX,
                const int cY)
{
    //ap_fixed<64,52> para_w = ap_fixed<64,52>(para_w_i);
    //ap_fixed<64,52> para_h = ap_fixed<64,52>(para_h_i);
    //ap_fixed<64,52> cX = ap_fixed<64,52>(cX_i);
    //ap_fixed<64,52> cY = ap_fixed<64,52>(cY_i);
    int f_1 = 65536; //1
    int f_1_5 = 98304; //1.5
    int f_2 = 131072;//2
    int f_4 = 262144;//4
    int maxIter = 255;
    __attribute__((xcl_pipeline_loop))
    for (int x = 0; x < width; x++) {
        int thisx = (int)(x*f_1);
        int tmp, zx, zy;
        __attribute__((xcl_pipeline_loop))
        for( int y = 0; y < height; y++) {
            int thisy = (int)(y*f_1);
            int tm = mul(thisx, para_w);
            tm = tm - f_1;
            zx = mul(f_1_5,tm);//f_1_5.Mul(thisx.Mul(para_w) - f_1)
            zy = mul(thisy,para_h) - f_1;
            int i = maxIter;

            int zx2 = mul(zx,zx);
            int zy2 = mul(zy,zy);
            __attribute__((xcl_pipeline_loop))
            while ((zx2+zy2) < f_4 && i > 0){
                //tmp = zx.Mul(zx) - zy.Mul(zy) - fixed.Int52_12(2867)
                tmp = zx2 - zy2 + cX;
                int tm1 = mul(zx,zy);
                zy = mul(f_2,tm1) + cY;
                zx = tmp;
                i--;
                zx2 = mul(zx,zx);
                zy2 = mul(zy,zy);
            }
            int index = x*height+y;
            c[index]= i;
        }
    }
}