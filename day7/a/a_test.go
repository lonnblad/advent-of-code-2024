package a_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code-2024/day7/a"
	"github.com/stretchr/testify/require"
)

//go:embed a_input.example.txt
var exampleInputA string

const expectedExampleOutputA = 3749

func Test_A_Example(t *testing.T) {
	actualOutputA, actualErr := a.Run(exampleInputA)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutputA, actualOutputA)
}

//go:embed a_input.txt
var inputA string

const expectedOutputA = 1985268524462

func Test_A_Input(t *testing.T) {
	actualOutputA, actualErr := a.Run(inputA)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutputA, actualOutputA)
}
