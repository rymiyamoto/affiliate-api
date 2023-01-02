package repository

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("GO_ENV", "test")
	m.Run()
}
