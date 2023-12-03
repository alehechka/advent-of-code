package utils_test

import (
	"advent/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Indices(t *testing.T) {
	indices := utils.Indices("this is his story is", "is")
	assert.Equal(t, 4, len(indices))
	assert.Equal(t, 2, indices[0])
	assert.Equal(t, 5, indices[1])
	assert.Equal(t, 9, indices[2])
	assert.Equal(t, 18, indices[3])
}
