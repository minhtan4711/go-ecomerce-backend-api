package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), input %d, output %d", input, output, actual)
	// }

	assert.Equal(t, AddOne(2), 4, "AddOne(2) should be 3")
	assert.Nil(t, nil, nil)
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Print("not execute")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Print("executing")
}
