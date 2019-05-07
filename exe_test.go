package exe

import "testing"
import "github.com/stretchr/testify/assert"

func TestEcho(t *testing.T) {
	t.Parallel()
	output, err := New().Run(`echo foo bar`)
	assert.NoError(t, err)
	assert.Len(t, output, 1)
	assert.Equal(t, "foo bar", output[0])
}

func TestParentEnv(t *testing.T) {
	t.Parallel()
	output, err := New().Run("env")
	assert.NoError(t, err)
	assert.NotEmpty(t, output[0])
}
