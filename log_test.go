package log

import (
	"testing"
)

func TestName(t *testing.T) {
	log := NewLogger()
	log.Debug("dd")

}
