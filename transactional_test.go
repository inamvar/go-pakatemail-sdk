package pakat_test

import (
	"testing"

	pakat "github.com/inamvar/go-pakatemail-sdk"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactional(t *testing.T) {
	msg := &pakat.Message{
		To:          []pakat.Address{{Name: "<name>", Email: "user@example.com"}},
		Sender:      &pakat.Address{Name: "<sender name>", Email: "no-reply@yourdomain.com"},
		Subject:     "test",
		TextContent: "this is a test",
	}

	err := pakat.SendTransactional(msg)
	assert.Nil(t, err)
}
