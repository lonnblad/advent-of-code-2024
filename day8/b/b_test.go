package b_test

import (
	_ "embed"
	"testing"

	"github.com/lonnblad/advent-of-code-2024/day8/b"
	"github.com/stretchr/testify/require"
)

//go:embed b_input.example.txt
var exampleInputB string

const expectedExampleOutputB = 34

func Test_B_Example(t *testing.T) {
	actualOutputB, actualErr := b.Run(exampleInputB)
	require.NoError(t, actualErr)
	require.Equal(t, expectedExampleOutputB, actualOutputB)
}

//go:embed b_input.txt
var inputB string

const expectedOutputB = 813

func Test_B_Input(t *testing.T) {
	actualOutputB, actualErr := b.Run(inputB)
	require.NoError(t, actualErr)
	require.Equal(t, expectedOutputB, actualOutputB)
}
