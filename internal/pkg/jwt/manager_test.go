package jwt

import (
	"github.com/jxlwqq/blog-microservices/internal/pkg/config"
	"github.com/jxlwqq/blog-microservices/internal/pkg/log"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJWTManager_Generate(t *testing.T) {
	t.Parallel()
	logger := log.New()
	conf := &config.Config{
		JWT: config.JWT{
			Secret:  "secret",
			Expires: 3600,
		},
	}
	jwtManager := NewJWTManager(logger, conf)
	id := uint64(1)
	username := "jack"
	tokenStr, err := jwtManager.Generate(id, username)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
}

func TestJWTManager_Verify(t *testing.T) {
	t.Parallel()
	logger := log.New()
	conf := &config.Config{
		JWT: config.JWT{
			Secret:  "secret",
			Expires: 3600,
		},
	}
	jwtManager := NewJWTManager(logger, conf)
	id := uint64(2)
	username := "rose"
	tokenStr, err := jwtManager.Generate(id, username)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
	claims, err := jwtManager.Verify(tokenStr)
	require.NoError(t, err)
	require.Equal(t, id, claims.ID)
	require.Equal(t, username, claims.Username)
}