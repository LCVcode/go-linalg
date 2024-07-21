package matrix

import (
    "reflect"
    "testing"
)

func TestNewZeroMatrix(t *testing.T) {
    rows, cols := 3, 4
    matrix, err := NewZeroMatrix(rows, cols)

    if err != nil {
        panic(err)
    }

    if matrix.Rows != rows || matrix.Cols != cols {
        t.Fatalf("expected dimensions (%d, %d), got (%d, %d)", rows, cols, matrix.Rows, matrix.Cols)
    }
}

func TestNewMatrix(t *testing.T) {
    data := [][]float64{
        {1, 2, 3},
        {4, 5, 6},
    }
    matrix, err := NewMatrix(2, 3, data)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if !reflect.DeepEqual(matrix.Data, data) {
        t.Fatalf("expected %v, got %v", data, matrix.Data)
    }

    _, err = NewMatrix(0, 0, [][]float64{})
    if err == nil {
        t.Fatal("expected error for empty data, but got none")
    }

    _, err = NewMatrix(0, 0, [][]float64{{1, 2}, {3}})
    if err == nil {
        t.Fatal("expected error for rows with different lengths, but got none")
    }
}

func TestAdd(t *testing.T) {
    a := Matrix{
        Rows: 2,
        Cols: 2,
        Data: [][]float64{
            {1, 2},
            {3, 4},
        },
    }
    b := Matrix{
        Rows: 2,
        Cols: 2,
        Data: [][]float64{
            {5, 6},
            {7, 8},
        },
    }
    expected := Matrix{
        Rows: 2,
        Cols: 2,
        Data: [][]float64{
            {6, 8},
            {10, 12},
        },
    }

    result, err := a.Add(b)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if !reflect.DeepEqual(result.Data, expected.Data) {
        t.Fatalf("expected %v, got %v", expected.Data, result.Data)
    }

    c := Matrix{
        Rows: 3,
        Cols: 2,
        Data: [][]float64{
            {1, 2},
            {3, 4},
            {5, 6},
        },
    }

    _, err = a.Add(c)
    if err == nil {
        t.Fatal("expected error for matrices with different dimensions, but got none")
    }
}

func TestMultiply(t *testing.T) {
    a := Matrix{
        Rows: 2,
        Cols: 3,
        Data: [][]float64{
            {1, 2, 3},
            {4, 5, 6},
        },
    }
    b := Matrix{
        Rows: 3,
        Cols: 2,
        Data: [][]float64{
            {7, 8},
            {9, 10},
            {11, 12},
        },
    }
    expected := Matrix{
        Rows: 2,
        Cols: 2,
        Data: [][]float64{
            {58, 64},
            {139, 154},
        },
    }

    result, err := a.Multiply(b)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if !reflect.DeepEqual(result.Data, expected.Data) {
        t.Fatalf("expected %v, got %v", expected.Data, result.Data)
    }

    c := Matrix{
        Rows: 2,
        Cols: 2,
        Data: [][]float64{
            {1, 2},
            {3, 4},
        },
    }

    _, err = a.Multiply(c)
    if err == nil {
        t.Fatal("expected error for matrices with incompatible dimensions, but got none")
    }
}

func TestTranspose(t *testing.T) {
    a := Matrix{
        Rows: 2,
        Cols: 2,
        Data: [][]float64{
            {1, 2},
            {3, 4},
        },
    }
    expected := Matrix{
                Rows: 2,
                Cols: 2,
                Data: [][]float64{
                    {1, 3},
                    {2, 4},
        },
    }

    transpose := a.T()

    if !reflect.DeepEqual(transpose.Data, expected.Data) {
        t.Fatalf("expected %v, got %v", expected.Data, transpose.Data)
    }
}

// TestNewIdentityMatrix tests the creation of an identity matrix.
func TestNewIdentityMatrix(t *testing.T) {
    size := 3
    identity, err := NewIdentityMatrix(size)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    expected := Matrix{
        Rows: size,
        Cols: size,
        Data: [][]float64{
            {1, 0, 0},
            {0, 1, 0},
            {0, 0, 1},
        },
    }

    if !reflect.DeepEqual(identity, expected) {
        t.Fatalf("expected %v, got %v", expected, identity)
    }
}

// TestMatrixMultiplyIdentity tests matrix multiplication with an identity matrix.
func TestMatrixMultiplyIdentity(t *testing.T) {
    a := Matrix{
        Rows: 2,
        Cols: 2,
        Data: [][]float64{
            {1, 2},
            {3, 4},
        },
    }

    identity, err := NewIdentityMatrix(2)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    result, err := a.Multiply(identity)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if !reflect.DeepEqual(result, a) {
        t.Fatalf("expected %v, got %v", a, result)
    }
}

// TestNewRandomMatrix tests the creation of a matrix with random values.
func TestNewRandomMatrix(t *testing.T) {
    rows, cols := 3, 3
    min, max := 0.0, 10.0
    m, err := NewRandomMatrix(rows, cols, min, max)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if m.Rows != rows || m.Cols != cols {
        t.Fatalf("expected matrix of size %dx%d, got %dx%d", rows, cols, m.Rows, m.Cols)
    }

    for i := range m.Data {
        for j := range m.Data[i] {
            if m.Data[i][j] < min || m.Data[i][j] > max {
                t.Fatalf("value %f out of range [%f, %f]", m.Data[i][j], min, max)
            }
        }
    }
}
