package a_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code-2024/dayN/a"
	"github.com/stretchr/testify/require"
)

//go:embed a_input.example.txt
var exampleInputA string

const expectedExampleOutputA = 0

func Test_A_Example(t *testing.T) {
	actualOutputA, actualErr := a.Run(exampleInputA)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutputA, actualOutputA)
}

//go:embed a_input.txt
var inputA string

const expectedOutputA = 0

func Test_A_Input(t *testing.T) {
	actualOutputA, actualErr := a.Run(inputA)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutputA, actualOutputA)
}