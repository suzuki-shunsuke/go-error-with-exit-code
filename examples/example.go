package examples

import (
	"errors"
	"fmt"

	"github.com/suzuki-shunsuke/go-error-with-exit-code"
)

func createUser() error {
	err := errors.New("user name is required")
	return ecerror.Wrap(err, 100)
}

func Example() {
	err := createUser()
	fmt.Printf("The exit code is %d\n", ecerror.GetExitCode(err))
	// Output:
	// The exit code is 100
}
