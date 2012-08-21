
// Copyright (c) Harri Rautila, 2012

// This file is part of go.opt/matrix package. It is free software, distributed
// under the terms of GNU Lesser General Public License Version 3, or any later
// version. See the COPYING tile included in this archive.

// Package matrix implements column major matrices.
package matrix


// Minimal interface for linear algebra packages BLAS/LAPACK
type Matrix interface {
	// The number of rows in this matrix.
	Rows() int
	// The number of columns in this matrix.
	Cols() int
	// The number of elements in this matrix.
	NumElements() int
	// Returns underlying float64 array for BLAS/LAPACK routines. Returns nil
	// if matrix is complex128 valued.
	FloatArray() []float64
	// Returns underlying complex128 array for BLAS/LAPACK routines. Returns nil
	// if matrix is float64 valued matrix.
	ComplexArray() []complex128
	// Returns true if matrix is complex valued. False otherwise.
	IsComplex() bool
	// For all float valued matrices return the value of A[0,0]. Returns NaN
	// if not float valued.
	Float() float64
	// For all complex valued matrices return the value of A[0,0]. Returns
	// NaN if not complex valued.
	Complex() complex128
	// Matrix in string format.
	String() string
	// Make a copy  and return as Matrix interface type.
	MakeCopy() Matrix
	// Match size. Return true if equal.
	SizeMatch(int, int) bool
	// Get matrix size. Return pair (rows, cols).
	Size() (int, int)
	// Test for type equality.
	EqualTypes(...Matrix) bool
}

// Interface for real and complex scalars.
type Scalar interface {
	Float() float64
	Complex() complex128
}

// Float constant
type FScalar float64

// Return self
func (self FScalar) Float() float64        { return float64(self) }
// Return complex(self, 0)
func (self FScalar) Complex() complex128   { return complex(float64(self), 0) }
func (self FScalar) Add(a float64) FScalar { return FScalar(float64(self)+a) }
func (self FScalar) Scale(a float64) FScalar{ return FScalar(float64(self)*a) }
func (self FScalar) Inv() FScalar          { return FScalar(1.0/float64(self)) }

// Return real(self).
type CScalar complex128
func (self CScalar) Float() float64 { return float64(real(self)) }
// Return self.
func (self CScalar) Complex() complex128  { return complex128(self) }

// Stacking direction for matrix constructor.
type Stacking int
const StackDown = Stacking(0)
const StackRight = Stacking(1)

// Matrix constructor data order
type DataOrder int
const RowOrder = DataOrder(0)
const ColumnOrder = DataOrder(1)

// Matrix dimensions, rows, cols and leading index. For column major matrices 
// leading index is equal to row count.
type dimensions struct {
	rows int
	cols int
	// actual offset between leading index
	step int
}

// Return number of rows.
func (A *dimensions) Rows() int {
	if A == nil {
		return 0
	}
	return A.rows
}

// Return number of columns.
func (A *dimensions) Cols() int {
	if A == nil {
		return 0
	}
	return A.cols
}

// Return size of the matrix as rows, cols pair.
func (A *dimensions) Size() (int, int) {
	if A == nil {
		return 0, 0
	}
	return A.rows, A.cols
}

// Set dimensions. Does not affect element allocations.
func (A *dimensions) SetSize(nrows, ncols int) {
	A.rows = nrows
	A.cols = ncols
	A.step = A.rows
}

// Return the leading index size. Column major matrices it is row count.
func (A *dimensions) LeadingIndex() int {
	return A.step
}

// Return total number of elements.
func (A *dimensions) NumElements() int {
	if A == nil {
		return 0
	}
	return A.rows * A.cols
}


// Return true if size of A is equal to parameter size (rows, cols).
func (A *dimensions) SizeMatch(rows, cols int) bool {
	return A != nil && A.rows == rows && A.cols == cols
}

// Change matrix shape if number of elements match to rows*cols.
func Reshape(m Matrix, rows, cols int) {
	if rows*cols == m.NumElements() {
		switch m.(type) {
		case *FloatMatrix:
			m.(*FloatMatrix).SetSize(rows, cols)
		case *ComplexMatrix:
			m.(*ComplexMatrix).SetSize(rows, cols)
		}
	}
}
		
// Set x = y ie. copy y to x. Matrices must have same number of elements but are
// required to have same shape.
func Set(x, y Matrix) {
	if x.NumElements() != y.NumElements() {
		return
	}
	if ! x.EqualTypes(y) {
		return
	}
	switch x.(type) {
	case *FloatMatrix:
		copy(x.FloatArray(), y.FloatArray())
	case *ComplexMatrix:
		copy(x.ComplexArray(), y.ComplexArray())
	}
}
	
// Create a set of indexes from start to end-1 with interval step.
func MakeIndexSet(start, end, step int) []int {
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if end-start == 0 {
		return make([]int, 0, 1)
	}
	if step < 0 {
		step = 1
	}
	//sz := (end-start)/step + 1
	inds := make([]int, 0)
	for k := start; k < end; k += step {
		inds = append(inds, k)
	}
	return inds
}

// Create index set to access diagonal entries of matrix.
//  indexes := MakeDiagonalSet(A.Size())
func MakeDiagonalSet(rows, cols int) []int {
	if rows != cols {
		return []int{}
	}
	inds := make([]int, rows)
	for i :=0; i < rows; i++  {
		inds[i] = i*rows + i
	}
	return inds
}

// Create index set for a row in matrix M. 
func RowIndexes(m Matrix, row int) []int {
	nrows, N := m.Size()
	if row > nrows {
		return []int{}
	}
	iset := make([]int, N)
	for i := 0; i < N; i++ {
		k := (row + i) * m.Cols()
		iset[i] = k
	}
	return iset
}

// Create index set for a column in matrix M. 
func ColumnIndexes(m Matrix, col int) []int {
	N, ncols := m.Size()
	if col > ncols {
		return []int{}
	}
	iset := make([]int, N)
	for i := 0; i < N; i++ {
		k := col * N + i
		iset[i] = k
	}
	return iset
}

// Create index set for diagonal in matrix M. 
func DiagonalIndexes(m Matrix) []int {
	if m.Cols() != m.Rows() {
		return []int{}
	}
	ind := 0
	iset := make([]int, m.Rows())
	for i := 0; i < m.Rows(); i++ {
		iset[i] = ind
		ind += m.Rows()
	}
	return iset
}

// Local Variables:
// tab-width: 4
// End:
