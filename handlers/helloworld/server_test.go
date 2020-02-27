package helloworld

import (
	"testing"

	"github.com/fguy/helloworld-go/config"
	"github.com/stretchr/testify/assert"

	"go.uber.org/zap"
)

func TestNewHTTPServer(t *testing.T) {
	t.Parallel()

	server := NewHTTPServer(&config.AppConfig{}, zap.NewNop(), nil)
	assert.NotNil(t, server)
}
