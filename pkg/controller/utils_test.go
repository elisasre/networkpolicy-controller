package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		array    []string
		word     string
		expected bool
	}{
		{
			name:     "First match",
			array:    []string{"a", "b", "c"},
			word:     "a",
			expected: true,
		},
		{
			name:     "Last match",
			array:    []string{"a", "b", "c"},
			word:     "c",
			expected: true,
		},
		{
			name:     "No match",
			array:    []string{"a", "b", "c"},
			word:     "d",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Contains(tt.array, tt.word)
			assert.Equal(t, tt.expected, got)
		})
	}
}
