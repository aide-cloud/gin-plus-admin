package gerrors

import "testing"

func TestNew(t *testing.T) {
	var err error
	err = New(WithMessage("test"))
	t.Log(err)
}
