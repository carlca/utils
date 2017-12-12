package essentials

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// CheckError is function to check error with call stack
func CheckError(msg string, err error, both bool) {
	if both {
		if err == nil {
			fmt.Println(fmt.Sprintf("%v succeeded", msg))
		} else {
			log.Fatal(fmt.Sprintf("%v failed: %+v", msg, errors.WithStack(err)))
		}
	} else {
		if err != nil {
			log.Fatal(fmt.Sprintf("%v failed: %+v", msg, errors.WithStack(err)))
		}
	}
}
