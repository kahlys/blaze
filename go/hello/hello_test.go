package hello

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	tests := map[string]struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := Hello(tt.name)
			assert.Equal(t, got, tt.want)
		})
	}
}
