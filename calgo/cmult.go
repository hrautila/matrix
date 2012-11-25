
package calgo

/*
 // S : start column
 // M : rows in C
 // N : end column
 // P : rows in B
 //
void mat_mult(int S, int N, int M, int P, double *C, double *A, double *B) {
   int i, j, k, cc, cr, ar, br;
   double beta;
   cc = S*M;
   for (j = S; j < N; j++) {
     br = 0;
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

*/
// #cgo CFLAGS: -O3
import "C"
import "unsafe"


// Calculate matrix-matrix product C = A*B for column major ordered matrices.
// S is start column and N is end column in C, M is rows in A and P is rows in B.
// For full matrix product call with S=0, N=C.Cols.
func Mult(S, N, M, P int, C, A, B []float64) {
	C.mat_mult(C.int(S), C.int(M), C.int(N), C.int(P),
		(*C.double)(unsafe.Pointer(&C[0])),
		(*C.double)(unsafe.Pointer(&A[0])),
		(*C.double)(unsafe.Pointer(&B[0])))
}

// Local Variables:
// tab-width: 4
// indent-tabs-mode: nil
// End:

