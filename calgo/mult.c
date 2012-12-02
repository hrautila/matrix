
// Copyright (c) Harri Rautila, 2012

// This file is part of github.com/hrautila/matrix package. It is free software,
// distributed under the terms of GNU Lesser General Public License Version 3, or
// any later version. See the COPYING tile included in this archive.


// S : start column
// M : rows in C
// N : end column
// P : rows in B
//
void mat_mult(int S, int N, int M, int P, double *C, double *A, double *B) {
   int i, j, k, cc, cr, ar, br;
   double beta;
   cc = S*M;
   br = S*P;
   for (j = S; j < N; j++) {
     ar = 0;
     for (k = 0; k < P; k++) {
       // move C index to first row in current column 
       cr = cc;
       // beta is value of B[k,j] 
       beta = B[br];
       // zero in B[k,j] does not increment value in C[:,j] 
       if (beta != 0.0) {
	 // C[:,j] += A[:,k]*B[k,j] 
	 for (i = 0; i < M; i++) {
	   // update target cell 
	   C[cr] += A[ar]*beta;
	   // move to next row in memory order 
	   cr++;
	   ar++;
	 }
       } else {
	 // we skipped all rows in this column, move to start of next column 
	 ar += M;
       }
       // move to next row in B, here ar points to first row of next column in A
       br++;
     }
     // forward to start of next column in C 
     cc += M;
   }
}
