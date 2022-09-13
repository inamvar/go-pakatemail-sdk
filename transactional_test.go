package pakat_test

import (
	"testing"

	pakat "github.com/inamvar/go-pakatemail-sdk"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransactional(t *testing.T) {
	msg := &pakat.Message{
		To:          []pakat.Address{{Name: "Iman Namvar", Email: "iman.namvar@gmail.com"}},
		Sender:      &pakat.Address{Name: "keypad", Email: "no-reply@keypad.ir"},
		Subject:     "test",
		TextContent: "this is a test",
	}

	err := pakat.SendTransactional(msg)
	assert.Nil(t, err)
}
