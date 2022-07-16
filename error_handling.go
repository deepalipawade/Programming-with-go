// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func main() {

	// ParseInt function will return error because empty string cannot be parsed into an integer value.
	// val, err := strconv.ParseInt("132", 10,64) // this will give us err -> nil and val -> 132

	_, err := strconv.ParseInt("", 10, 64)
	if err != nil {

		fmt.Printf("\n Actual error : %+v \n", err)
		// Wrap or Wrapf helps to identify which function failed by adding(wrapping) this extra string msg to the error.

		err = errors.Wrap(err, "parse int failed")
		// return err ......usually we return from here as error has occured
	}

	// error in golang are just values so you can print them or use conditional operators on them , see line 14 above where we checked if err is nil.
	fmt.Println("\n\n ...... Wrapped error without trace ==>>>> ", err)

	// Using %+v print entire stack trace of error along with its value
	fmt.Printf("\n--------------- wrapped error with stack trace ==>>> %+v\n", err)

	// Cause method is used to umwrap error which reverse of Wrap (line 17) this line will print actual err if we want to check for Sentinal Error codes.
	fmt.Printf("\n--------------- unwrapped error ==>>> %+v \n", errors.Cause(err))

}
