package day4_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code-2024/day4"
	"github.com/stretchr/testify/require"
)

//go:embed a_input.example.txt
var exampleInputA string

const expectedExampleOutputA = 18

func Test_A_Example(t *testing.T) {
	actualOutputA, actualErr := day4.A(exampleInputA)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutputA, actualOutputA)
}

//go:embed a_input.txt
var inputA string

const expectedOutputA = 2639

func Test_A_Input(t *testing.T) {
	actualOutputA, actualErr := day4.A(inputA)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutputA, actualOutputA)
}
