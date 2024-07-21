package main

import (
    "github.com/LCVcode/linalg/matrix"
)

func main() {
    matrix, err := matrix.NewRandomMatrix(10, 10, -1, 1)
    if err != nil {
        panic(err)
    }

    matrix.Display(4)
}
