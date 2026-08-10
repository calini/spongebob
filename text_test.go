package spongebob

import (
	"strings"
	"testing"
)

// TestText checks that transformation is happening without ruining context
func TestText(t *testing.T) {
	input := "Hello world!"
	result := Text(input)

	if !strings.EqualFold(input, result) {
		t.Errorf("Should have same content except capitalisation: \"%s\" - \"%s\"", input, result)
	}
}
