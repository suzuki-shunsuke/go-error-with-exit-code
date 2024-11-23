package ecerror

import (
	"errors"
	"testing"
)

func TestWrap(t *testing.T) {
	t.Parallel()
	baseErr := errors.New("hello")
	err := Wrap(baseErr, 1)
	var e withExitCodeError
	if errors.As(err, &e) {
		if !errors.Is(e.err, baseErr) {
			t.Fatal("err.err != baseErr")
		}
		return
	}
	t.Fatal("error must be withExitCodeError")
}

func Test_withExitCodeError_ExitCode(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		err   withExitCodeError
		exp   int
	}{
		{
			err: withExitCodeError{
				code: 5,
			},
			exp: 5,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			act := d.err.ExitCode()
			if act != d.exp {
				t.Fatalf("err.ExitCode() got %d, want %d", act, d.exp)
			}
		})
	}
}

func Test_withExitCodeError_Error(t *testing.T) {
	t.Parallel()
	data := []struct {
		title string
		err   withExitCodeError
		exp   string
	}{
		{
			title: "normal",
			err: withExitCodeError{
				err: errors.New("hello"),
			},
			exp: "hello",
		},
		{
			title: "err.err is nil",
			err: withExitCodeError{
				err: nil,
			},
			exp: "",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			act := d.err.Error()
			if act != d.exp {
				t.Fatalf("err.Error() got %s, want %s", act, d.exp)
			}
		})
	}
}

func Test_withExitCodeError_Unwrap(t *testing.T) {
	t.Parallel()
	helloError := errors.New("hello")
	data := []struct {
		title string
		err   withExitCodeError
		exp   error
	}{
		{
			title: "normal",
			err: withExitCodeError{
				err: helloError,
			},
			exp: helloError,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			act := d.err.Unwrap()
			if !errors.Is(act, d.exp) {
				t.Fatalf("err.Unwrap() got %v, want %v", act, d.exp)
			}
		})
	}
}

func TestGetExitCode(t *testing.T) {
	t.Parallel()
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
			err: withExitCodeError{
				code: 5,
			},
			exp: 5,
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			t.Parallel()
			act := GetExitCode(d.err)
			if act != d.exp {
				t.Fatalf("GetExitCode(err) got %d, want %d", act, d.exp)
			}
		})
	}
}
