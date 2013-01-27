// Copyright (c) Harri Rautila, 2012

// This file is part of go.opt/matrix package. It is free software, distributed
// under the terms of GNU Lesser General Public License Version 3, or any later
// version. See the COPYING tile included in this archive.

package matrix

// Test for equality. Return true if for all i,j: all A[i,j] = B[i,j]
func (A *ComplexMatrix) Equal(B *ComplexMatrix) bool {
    if A.Rows() != B.Rows() || A.Cols() != B.Cols() {
        return false
    }
	nrows := A.Rows()
	step := A.LeadingIndex()
	for k := 0; k < A.NumElements(); k++ {
		rk := realIndex(k, nrows, step)
		if A.elements[rk] != B.elements[rk] {
			return false
		}
	}
    return true
}

// Local Variables:
// tab-width: 4
// End:
