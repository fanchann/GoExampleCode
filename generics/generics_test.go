package generics_test

import (
	"go-example-code/generics"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintToTerminal(t *testing.T) {
	gen := generics.GenericsFunc("Generics is awesome")
	assert.NotNil(t, gen)
}
