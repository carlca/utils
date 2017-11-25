package essentials

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// CheckError is function to check error with call stack
func CheckError(err error) {
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
		os.Exit(1) // or anything else ...
	}
}
