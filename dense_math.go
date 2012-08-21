
// Copyright (c) Harri Rautila, 2012

// This file is part of go.opt/matrix package. It is free software, distributed
// under the terms of GNU Lesser General Public License Version 3, or any later
// version. See the COPYING tile included in this archive.

package matrix

import "math"


// Compute in-place A *= alpha for all elements in the matrix if list of indexes
// is empty. Otherwise compute A[i] *= alpha for indexes in column-major order.
func (A *FloatMatrix) Scale(alpha float64, indexes... int) *FloatMatrix {
	if len(indexes) == 0 {
		for k, _ := range A.elements {
			A.elements[k] *= alpha
		}
	} else  {
		N := A.NumElements()
		for _, k := range indexes {
			if k < 0 {
				k = N + k
			}
			A.elements[k] *= alpha
		}
	}
	return A
}

// Compute in-place A[indexes[i]] *= values[i]. Indexes are in column-major order.
func (A *FloatMatrix) ScaleIndexes(indexes []int, values []float64) *FloatMatrix {
	if len(indexes) == 0 {
		return A
	}
	N := A.NumElements()
	for i, k := range indexes {
		if i >= len(values) {
			return A
		}
		if k < 0 {
			k = N + k
		}
		A.elements[k] *= values[i]
	}
	return A
}


// Compute in-place remainder A[i,j] %= alpha
func (A *FloatMatrix) Mod(alpha float64, indexes... int) *FloatMatrix {
	if len(indexes) == 0 {
		for k, _ := range A.elements {
			A.elements[k] = math.Mod(A.elements[k], alpha)
		}
	} else  {
		N := A.NumElements()
		for _, k := range indexes {
			if k < 0 {
				k = N + k
			}
			A.elements[k] = math.Mod(A.elements[k], alpha)
		}
	}
	return A
}

// Compute in-place A += alpha for all elements in the matrix if list of indexes
// is empty. Otherwise compute A[i] += alpha for indexes in column-major order.
func (A *FloatMatrix) Add(alpha float64, indexes... int) *FloatMatrix {
	if len(indexes) == 0 {
		for k, _ := range A.elements {
			A.elements[k] += alpha
		}
	} else  {
		N := A.NumElements()
		for _, k := range indexes {
			if k < 0 {
				k = N + k
			}
			A.elements[k] += alpha
		}
	}
	return A
}

// Compute in-place A[indexes[i]] += values[i]. Indexes are in column-major order.
func (A *FloatMatrix) AddIndexes(indexes []int, values []float64) *FloatMatrix {
	if len(indexes) == 0 {
		return A
	}
	N := A.NumElements()
	for i, k := range indexes {
		if i >= len(values) {
			return A
		}
		if k < 0 {
			k = N + k
		}
		A.elements[k] += values[i]
	}
	return A
}


// Compute element wise division C[i,] = A[i,j] / B[i,j]. Returns new matrix.
func (A *FloatMatrix) Div(B *FloatMatrix) *FloatMatrix {
	if ! A.SizeMatch(B.Size()) {
		return nil
	}
	C := FloatZeros(A.Rows(), A.Cols())
	for k, v := range B.elements {
		C.elements[k] = A.elements[k] / v
	}
	return C
}

// Compute element-wise product C[i,j] = A[i,j] * B[i,j]. Returns new matrix.
func (A *FloatMatrix) Mul(B *FloatMatrix) *FloatMatrix {
	if ! A.SizeMatch(B.Size()) {
		return nil
	}
	var C *FloatMatrix = FloatZeros(A.Rows(), A.Cols())
	for k, _ := range B.elements {
		val := A.elements[k] * B.elements[k]
		C.elements[k] = val
	}
	return C
}

// Compute element-wise sum C = A + B. Returns a new matrix.
func (A *FloatMatrix) Plus(B *FloatMatrix) *FloatMatrix {
	if ! A.SizeMatch(B.Size()) {
		return nil
	}
	C := FloatZeros(A.Rows(), A.Cols())
	for k, _ := range A.elements {
		C.elements[k] = A.elements[k] + B.elements[k]
	}
	return C
}

// Compute element-wise difference C = A - B. Returns a new matrix.
func (A *FloatMatrix) Minus(B *FloatMatrix) *FloatMatrix {
	if ! A.SizeMatch(B.Size()) {
		return nil
	}
	C := FloatZeros(A.Rows(), A.Cols())
	for k, _ := range A.elements {
		C.elements[k] = A.elements[k] - B.elements[k]
	}
	return C
}


// Compute matrix product C = A * B where A is m*p and B is p*n.
// Returns a new m*n matrix.
func (A *FloatMatrix) Times(B *FloatMatrix) *FloatMatrix {
	if A.Cols() != B.Rows() {
		return nil
	}
	rows := A.Rows()
	cols := B.Cols()
	C := FloatZeros(rows, cols)
	arow := make([]float64, A.Cols())
	bcol := make([]float64, B.Rows())
	for i := 0; i < rows; i++ {
		arow = A.GetRowArray(i, arow)
		for j := 0; j < cols; j++ {
			bcol = B.GetColumnArray(j, bcol)
			for k, _ := range arow {
				C.elements[j*rows+i] += arow[k]*bcol[k]
			}
		}
	}
	return C
}


// Compute A = fn(C) by applying function fn element wise to C.
// For all i, j: A[i,j] = fn(C[i,j]). If C is nil then computes inplace
// A = fn(A). If C is not nil then sizes of A and C must match.
// Returns pointer to self.
func (A *FloatMatrix) Apply(C *FloatMatrix, fn func(float64)float64) *FloatMatrix {
	if C != nil && ! A.SizeMatch(C.Size()) {
		return nil
	}
	var B *FloatMatrix = C
	if B == nil {
		B = A
	}
	for k,v := range B.elements {
		A.elements[k] = fn(v)
	}
	return A
}

// Compute A = fn(C) by applying function fn to all elements in indexes.
// For all i in indexes: A[i] = fn(C[i]).
// If C is nil then computes inplace A = fn(A). If C is not nil then sizes of A and C must match.
// Returns pointer to self.
func (A *FloatMatrix) ApplyToIndexes(C *FloatMatrix, indexes []int, fn func(float64)float64) *FloatMatrix {
	if C != nil && ! A.SizeMatch(C.Size()) {
		return nil
	}
	B := C
	if C == nil {
		B = A
	}
	if len(indexes) > 0 {
		for _,v := range indexes {
			A.elements[v] = fn(B.elements[v])
		}
	}
	return A
}

// Compute A = fn(C, x) by applying function fn element wise to C.
// For all i, j: A[i,j] = fn(C[i,j], x). 
func (A *FloatMatrix) ApplyConst(C *FloatMatrix, fn func(float64,float64)float64, x float64) *FloatMatrix {

	if C != nil && ! A.SizeMatch(C.Size()) {
		return nil
	}
	B := C
	if C == nil {
		B = A
	}
	for k,v := range B.elements {
		A.elements[k] = fn(v, x)
	}
	return A
}

// Find element-wise maximum. 
func (A *FloatMatrix) Max(indexes... int) float64 {
	m := math.Inf(-1)
	if len(indexes) == 0 {
		for _, v := range A.elements {
			m = math.Max(m, v)
		}
	} else {
		N := A.NumElements()
		for _, k := range indexes {
			if k < 0 {
				k = N + k
			}
			m = math.Max(m, A.elements[k])
		}
	}
	return m
}

// Find element-wise minimum. 
func (A *FloatMatrix) Min(indexes... int) float64 {
	m := math.Inf(+1)
	if len(indexes) == 0 {
		for _, v := range A.elements {
			m = math.Min(m, v)
		}
	} else {
		N := A.NumElements()
		for _, k := range indexes {
			if k < 0 {
				k = N + k
			}
			m = math.Min(m, A.elements[k])
		}
	}
	return m
}

// Return sum of elements
func (A *FloatMatrix) Sum(indexes... int) float64 {
	m := 0.0
	if len(indexes) == 0 {
		for _, v := range A.elements {
			m += v
		}
	} else {
		N := A.NumElements()
		for _, k := range indexes {
			if k < 0 {
				k = N + k
			}
			m += A.elements[k]
		}
	}
	return m
}

// Compute element-wise C = Exp(A). Returns a new matrix.
func (A *FloatMatrix) Exp() *FloatMatrix {
	C := FloatZeros(A.Rows(), A.Cols())
	return C.Apply(A, math.Exp)
}

// Compute element-wise C = Log(A). Returns a new matrix.
func (A *FloatMatrix) Log() *FloatMatrix {
	C := FloatZeros(A.Rows(), A.Cols())
	return C.Apply(A, math.Log)
}
		
// Compute element-wise C = Pow(A). Returns a new matrix.
func (A *FloatMatrix) Pow(exp float64) *FloatMatrix {
	C := FloatZeros(A.Rows(), A.Cols())
	return C.ApplyConst(A, math.Pow, exp)
}
		

// Local Variables:
// tab-width: 4
// End:
