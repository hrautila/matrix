
package calgo

import (
    "github.com/hrautila/matrix"
    "testing"
)

const NROWS = 2000
var A, B, C *matrix.FloatMatrix

func TestMakeData(t *testing.T) {
    A = matrix.FloatWithValue(NROWS, NROWS, 2.0)
    B = matrix.FloatDiagonal(NROWS, 1.0)
    C = matrix.FloatZeros(NROWS, NROWS)
    t.Logf("A [%d,%d] non-zero matrix\n", NROWS, NROWS)
    t.Logf("B [%d,%d] diagonal matrix\n", NROWS, NROWS)
    t.Logf("C [%d,%d] result matrix\n", NROWS, NROWS)
    
}

func TestMultAB(t *testing.T) {
    Mult(0, C.Cols(), A.Rows(), B.Rows(), C.FloatArray(), A.FloatArray(), B.FloatArray())
}

func TestMultBA(t *testing.T) {
    Mult(0, C.Cols(), A.Rows(), B.Rows(), C.FloatArray(), B.FloatArray(), A.FloatArray())
}

// Local Variables:
// tab-width: 4
// indent-tabs-mode: nil
// End:
