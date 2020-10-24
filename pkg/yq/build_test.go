package yq

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMixin_Build(t *testing.T) {
	t.Run("default client version", func(t *testing.T) {
		m := NewTestMixin(t)

		err := m.Build()

		require.NoError(t, err, "Build failed")
		assert.Contains(t, m.TestContext.GetOutput(), "3.4.1")
	})

	t.Run("client version specified", func(t *testing.T) {
		m := NewTestMixin(t)

		mixinInputB, err := ioutil.ReadFile("testdata/build-input-with-client-version.yaml")
		require.NoError(t, err)
		m.In = bytes.NewBuffer(mixinInputB)

		err = m.Build()

		require.NoError(t, err, "Build failed")
		assert.Contains(t, m.TestContext.GetOutput(), "3.4.0")
	})
}
