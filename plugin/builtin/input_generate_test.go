package builtin

import (
	"testing"

	"github.com/bluemedora/bplogagent/plugin"
	"github.com/stretchr/testify/assert"
)

func TestGenerateImplementations(t *testing.T) {
	assert.Implements(t, (*plugin.Plugin)(nil), new(GenerateInput))
	assert.Implements(t, (*plugin.Producer)(nil), new(GenerateInput))
}
