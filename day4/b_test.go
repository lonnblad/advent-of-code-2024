package day4_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code-2024/day4"
	"github.com/stretchr/testify/require"
)

//go:embed b_input.example.txt
var exampleInputB string

const expectedExampleOutputB = 9

func Test_B_Example(t *testing.T) {
	actualOutputB, actualErr := day4.B(exampleInputB)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutputB, actualOutputB)
}

//go:embed b_input.txt
var inputB string

const expectedOutputB = 2005

func Test_B_Input(t *testing.T) {
	actualOutputB, actualErr := day4.B(inputB)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutputB, actualOutputB)
}
