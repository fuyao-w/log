package log

import (
	"path/filepath"
	"testing"
)

func TestName(t *testing.T) {
	dir := "/~a/b/c/d/e"
	t.Log(filepath.Base(dir))

}
