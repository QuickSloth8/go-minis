/*
A program that tests the effects of defer statements/functions
on returned function values. Also, it addresses the case
of errors rising during defer execution.
*/
package main

import "fmt"

// handle errors in defer internally
func foo1() (string, error) {
	defer func() {
		err := func() error {
			return fmt.Errorf("returned from defer")
		}()
		printWithFuncName("foo1", err.Error())
	}()

	printWithFuncName("foo1", "body")

	return "foo1() return", nil
}

// the same named returns in defer function
// note how it doesn't affect the named
// returns of the outer scope function
func foo2() (s string, err error) {
	defer func() (s string, err error) {
		return s, fmt.Errorf("error from defer")
	}()

	printWithFuncName("foo2", "body")

	s = "foo2() return"
	return s, nil
}

// demonstrate how the named returns can be changed in defer
func foo3() (s string, err error) {
	defer func() {
		err = fmt.Errorf("foo3() changed error")
	}()

	printWithFuncName("foo3", "body")

	s = "foo3() return"
	return s, nil
}

// demonstrates how it is not possible to change return values
// without named returns
func foo4() (string, error) {
	var s string
	var err error

	defer func() {
		s = "foo4() changed return"
		err = fmt.Errorf("foo4() changed error")
	}()

	printWithFuncName("foo4", "body")

	s = "foo4() return"

	defer func() {
		s = "foo4() changed return, again"
		err = fmt.Errorf("foo4() changed error, again")
	}()

	return s, nil
}

func main() {
	r, err := foo1()
	printWithFuncName("main", fmt.Sprintf("%+v, %+v\n", r, err))

	r, err = foo2()
	printWithFuncName("main", fmt.Sprintf("%+v, %+v\n", r, err))

	r, err = foo3()
	printWithFuncName("main", fmt.Sprintf("%+v, %+v\n", r, err))

	r, err = foo4()
	printWithFuncName("main", fmt.Sprintf("%+v, %+v\n", r, err))

}

func printWithFuncName(funcName, message string) {
	fmt.Println(funcName, "\t-", message)
}
