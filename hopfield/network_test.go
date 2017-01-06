package hopfield

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeStateNeuron(t *testing.T) {
	assert := assert.New(t)

	n := &Neuron{state: 1.0}
	assert.True(!n.ChangeState(1.0))
	assert.True(n.ChangeState(-1.0))
}

func TestNewnet(t *testing.T) {
	assert := assert.New(t)

	size := 5
	n, err := NewNet(size)
	assert.NotNil(n)
	assert.NoError(err)

	n, err = NewNet(-2)
	assert.Nil(n)
	assert.Error(err)
}

func TestWeights(t *testing.T) {
	assert := assert.New(t)

	size := 5
	n, err := NewNet(size)
	assert.NotNil(n)
	assert.NoError(err)

	w := n.Weights()
	rows, cols := w.Dims()
	assert.Equal(rows, size)
	assert.Equal(cols, size)
}

func TestNeurons(t *testing.T) {
	assert := assert.New(t)

	size := 5
	n, err := NewNet(size)
	assert.NotNil(n)
	assert.NoError(err)

	neurons := n.Neurons()
	assert.Equal(size, len(neurons))
}

func TestBias(t *testing.T) {
	assert := assert.New(t)

	size := 5
	n, err := NewNet(size)
	assert.NotNil(n)
	assert.NoError(err)

	bias := n.Bias()
	rows, cols := bias.Dims()
	assert.Equal(size, rows)
	assert.Equal(1, cols)
}

func TestStore(t *testing.T) {
	assert := assert.New(t)

	size := 4
	n, err := NewNet(size)
	assert.NotNil(n)
	assert.NoError(err)

	var pattern Pattern
	errString := "Invalid pattern supplied: %v\n"
	err = n.Store(pattern)
	assert.EqualError(err, fmt.Sprintf(errString, pattern))

	pattern = Pattern{1.0, -1.0}
	errString = "Dimension mismatch: %v\n"
	err = n.Store(pattern)
	assert.EqualError(err, fmt.Sprintf(errString, pattern))

	pattern = Pattern{1.0, -1.0, -1.0, 1.0}
	err = n.Store(pattern)
	assert.NoError(err)
	assert.Equal(n.Weights().At(0, 3), n.Weights().At(3, 0))
}

func TestRestore(t *testing.T) {
	assert := assert.New(t)

	size := 4
	maxiters := 10
	eqiters := 5
	n, err := NewNet(size)
	assert.NotNil(n)
	assert.NoError(err)

	pattern := Pattern{1.0, -1.0, -1.0, 1.0}
	err = n.Store(pattern)
	assert.NoError(err)

	pattern = Pattern(nil)
	errString := "Invalid pattern supplied: %v\n"
	res, err := n.Restore(pattern, maxiters, eqiters)
	assert.Nil(res)
	assert.EqualError(err, fmt.Sprintf(errString, pattern))

	pattern = Pattern{1.0, -1.0}
	errString = "Dimension mismatch: %v\n"
	res, err = n.Restore(pattern, maxiters, eqiters)
	assert.Nil(res)
	assert.EqualError(err, fmt.Sprintf(errString, pattern))

	maxiters = -5
	pattern = Pattern{1.0, -1.0, -1.0, 1.0}
	errString = "Invalid number of max iterations: %d\n"
	res, err = n.Restore(pattern, maxiters, eqiters)
	assert.Nil(res)
	assert.EqualError(err, fmt.Sprintf(errString, maxiters))
	maxiters = 10

	eqiters = -3
	errString = "Invalid number of equilibrium iterations: %d\n"
	res, err = n.Restore(pattern, maxiters, eqiters)
	assert.Nil(res)
	assert.EqualError(err, fmt.Sprintf(errString, eqiters))
	eqiters = 5

	res, err = n.Restore(pattern, maxiters, eqiters)
	assert.NotNil(res)
	assert.NoError(err)

	pattern = Pattern{-1.0, -1.0, 1.0, 1.0}
	res, err = n.Restore(pattern, maxiters, eqiters)
	assert.NotNil(res)
	assert.NoError(err)
}

func TestEnergy(t *testing.T) {
	assert := assert.New(t)

	size := 4
	n, err := NewNet(size)
	assert.NotNil(n)
	assert.NoError(err)

	pattern := Pattern{1.0, -1.0, -1.0, 1.0}
	energy, err := n.Energy(pattern)
	assert.Equal(0.0, energy)
	assert.NoError(err)

	pattern = Pattern(nil)
	errString := "Invalid pattern supplied: %v\n"
	energy, err = n.Energy(pattern)
	assert.Equal(0.0, energy)
	assert.EqualError(err, fmt.Sprintf(errString, pattern))

	pattern = Pattern{1.0, -1.0}
	errString = "Dimension mismatch: %v\n"
	energy, err = n.Energy(pattern)
	assert.Equal(0.0, energy)
	assert.EqualError(err, fmt.Sprintf(errString, pattern))
}
