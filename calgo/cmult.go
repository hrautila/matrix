
// Copyright (c) Harri Rautila, 2012

// This file is part of github.com/hrautila/matrix package. It is free software,
// distributed under the terms of GNU Lesser General Public License Version 3, or
// any later version. See the COPYING tile included in this archive.

package calgo

/*
extern void mat_mult(int S, int N, int M, int P, double *C, double *A, double *B);

*/
// #cgo CFLAGS: -O3
import "C"
import "unsafe"


// Calculate matrix-matrix product C = A*B for column major ordered matrices.
// S is start column and N is end column in C, M is rows in A and P is rows in B
// and columns in A. For full matrix product call with S=0, N=C.Cols.
func Mult(S, N, M, P int, C, A, B []float64) {
	C.mat_mult(C.int(S), C.int(N), C.int(M), C.int(P),
		(*C.double)(unsafe.Pointer(&C[0])),
		(*C.double)(unsafe.Pointer(&A[0])),
		(*C.double)(unsafe.Pointer(&B[0])))
}

// Local Variables:
// tab-width: 4
// indent-tabs-mode: nil
// End:

