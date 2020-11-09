package routes

import (
	"github.com/raulinoneto/atm-withdrawal-analisys/config/container"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetRoutes(t *testing.T) {
	assert.Len(t, GetRoutes(new(container.Container)), 1)
}