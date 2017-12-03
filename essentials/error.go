package essentials

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// CheckError is function to check error with call stack
func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%v: %+v", msg, errors.WithStack(err)))
	}
}
