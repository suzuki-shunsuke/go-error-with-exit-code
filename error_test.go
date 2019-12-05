package ecerror

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	baseErr := errors.New("hello")
	err := Wrap(baseErr, 1)
	e := err.(*withExitCodeError)
	if e.err != baseErr {
		t.Fatal("err.err != baseErr")
	}
}

func Test_withExitCodeError_ExitCode(t *testing.T) {
	err := &withExitCodeError{
		code: 5,
	}
	exp := 5
	act := err.ExitCode()
	if act != exp {
		t.Fatalf("err.ExitCode() got %d, want %d", act, exp)
	}
}

func Test_withExitCodeError_Error(t *testing.T) {
	exp := "hello"
	err := &withExitCodeError{
		err: errors.New(exp),
	}
	act := err.Error()
	if act != exp {
		t.Fatalf("err.Error() got %s, want %s", act, exp)
	}
}

func Test_withExitCodeError_Unwrap(t *testing.T) {
	exp := errors.New("hello")
	err := &withExitCodeError{
		err: exp,
	}
	act := err.Unwrap()
	if act != exp {
		t.Fatalf("err.Error() got %v, want %v", act, exp)
	}
}

func TestGetExitCode(t *testing.T) {
	data := []struct {
		title string
		err   error
		exp   int
	}{
		{
			title: "error is nil",
			err:   nil,
			exp:   0,
		},
		{
			title: "not withExitCodeError",
			err:   errors.New("hello"),
			exp:   1,
		},
		{
			title: "withExitCodeError",
			err: &withExitCodeError{
				code: 5,
			},
			exp: 5,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			act := GetExitCode(d.err)
			if act != d.exp {
				t.Fatalf("GetExitCode(err) got %d, want %d", act, d.exp)
			}
		})
	}
}
