package matrix

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

// Matrix represents a mathematical matrix
type Matrix struct {
    Rows int
    Cols int
    Data [][]float64
}

// Creates a new Matrix
// Returns an error if dimensions are not greater than 0 or data shape is mismatched
func NewMatrix(rows, cols int, data[][]float64) (Matrix, error) {
    if rows <= 0 || cols <= 0 {
        return Matrix{}, errors.New("dimensions must be positive integers")
    }
    if len(data) != rows {
        return Matrix{}, errors.New("invalid data dimension: rows")
    }
    if len(data[0]) != cols {
        return Matrix{}, errors.New("invalid data dimension: columns")
    }
    return Matrix{Rows: rows, Cols: cols, Data: data}, nil
}

// Creates a new Matrix initialized with zeroes
func NewZeroMatrix(rows, cols int) (Matrix, error) {
    data := make([][]float64, rows)
    for i := range data {
        data[i] = make([]float64, cols)
    }

    result, err := NewMatrix(rows, cols, data)

    if err != nil {
        panic(err)
    }

    return result, nil
}

// Creates a new identify Matrix of a given size
func NewIdentityMatrix(size int) (Matrix, error) {
    if size <= 0 {
        return Matrix{}, errors.New("Expected an identity matrix size greater than 0")
    }

    identity, err := NewZeroMatrix(size, size)

    if err != nil {
        panic(err)
    }

    for i := 0; i<size; i++ {
        identity.Data[i][i] = 1.0
    }

    return identity, nil
}

// Adds to matrices together
func (m Matrix) Add(other Matrix) (Matrix, error) {
    if m.Rows != other.Rows || m.Cols != other.Cols {
        return Matrix{}, errors.New("matrices must have matching dimensions")
    }
    result := make([][]float64, m.Rows)
    for i:= range m.Data {
        result[i] = make([]float64, m.Rows)
        for j := range m.Data[i] {
            result[i][j] = m.Data[i][j] + other.Data[i][j]
        }
    }
    return Matrix{
        Rows: m.Rows,
        Cols: m.Cols,
        Data: result,
    }, nil
}

// Multiple performs matrix multiplication between two matrices.
// Returns and error if matrices have incompatible dimensions.
func (m Matrix) Multiply(other Matrix) (Matrix, error) {
    if m.Cols != other.Rows {
        return Matrix{}, errors.New("incompatible dimensions for matrix multiplication")
    }

    result, err := NewZeroMatrix(m.Rows, other.Cols)

    if err != nil {
        panic(err)
    }

    for i := range m.Data {
        for j := range other.Data[0] {
            for k := range m.Data[0] {
                result.Data[i][j] += m.Data[i][k] * other.Data[k][j]
            }
        }
    }

    return result, nil
}

// Transpose returns the transpose of a matrix.
func (m Matrix) T() (Matrix) {
    transpose, err := NewZeroMatrix(m.Cols, m.Rows)

    if err != nil {
        panic(err)
    }

    for j := range m.Data {
        for i := range m.Data[0] {
            transpose.Data[i][j] = m.Data[j][i]
        }
    }

    return transpose
}

// Apply a function to all the elements in a matrix.
// The given function must take and return a float64
func (m Matrix) Map(f func(float64) float64) (Matrix, error) {
    result, err := NewZeroMatrix(m.Rows, m.Cols)

    if err != nil {
        panic(err)
    }

    for i := range m.Data {
        for j := range m.Data[0] {
            // TODO: add error handling here
            result.Data[i][j] = f(m.Data[i][j])
        }
    }

    return result, nil
}

// NewRandomMatrix creates a new matrix with random values between min and max.
func NewRandomMatrix(rows, cols int, min, max float64) (Matrix, error) {
    if rows <= 0 || cols <= 0 {
        return Matrix{}, errors.New("number of rows and columns must be greater than 0")
    }

    rand.Seed(time.Now().UnixNano())
    data := make([][]float64, rows)
    for i := range data {
        data[i] = make([]float64, cols)
        for j := range data[i] {
            data[i][j] = min + rand.Float64()*(max-min)
        }
    }

    return Matrix{Rows: rows, Cols: cols, Data: data}, nil
}

// calculateWidth is a helper function to calculate the largest absolute value in each column of a matrix.
// Used to align decimal places when displaying matrices.
func calculateWidth(data [][]float64, precision int) []int {
    if len(data) == 0 || len(data[0]) == 0 {
        return nil
    }

    cols := len(data[0])
    widths := make([]int, cols)

    for col := 0; col < cols; col++ {
        maxWidth := 0
        for row := range data {
            formatted := fmt.Sprintf("%.*f", precision, data[row][col])
            if len(formatted) > maxWidth {
                maxWidth = len(formatted)
            }
        }
        widths[col] = maxWidth
    }

    return widths
}

// Display prints the matrix in a formatted manner.
// Each element is rounded to the specified precision.
// Columns are aligned based on the maximum width of their elements.
// Square brackets are added to the start and end of each row.
//
// Arguments:
// - precision: the number of decimal places to which each element is rounded.
func (m Matrix) Display(precision int) {
    widths := calculateWidth(m.Data, precision)
    if widths == nil {
        fmt.Println("Empty matrix")
        return
    }

    for _, row := range m.Data {
        fmt.Print("[")
        for j, val := range row {
            format := fmt.Sprintf("%%%d.%df", widths[j], precision)
            fmt.Printf(format, val)
            if j < len(row)-1 {
                fmt.Print(" ")
            }
        }
        fmt.Println("]")
    }
}

