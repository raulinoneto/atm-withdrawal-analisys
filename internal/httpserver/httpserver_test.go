package httpserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestServer_Run(t *testing.T) {
	s := New(&Options{
		Host:        "4321",
	})
	go s.Run()
	time.Sleep(2 * time.Second)
	assert.Equal(t, true, s.IsRunning(), "Server is not running")
}
