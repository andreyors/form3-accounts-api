package form3

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	f3 := New("http://0.0.0.0:8080/", time.Duration(5*time.Second))

	assert.IsType(t, &Form3{}, f3)
}
