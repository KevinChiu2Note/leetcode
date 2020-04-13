package design_pattern

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAPI(t *testing.T) {
	NewAPI("en").Say()
	NewAPI("cn").Say()
	a3 := NewAPI("")
	assert.Equal(t, nil, a3)
}
