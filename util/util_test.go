package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func ExampleNewErrorPrefixer() {
	prefixer := NewErrorPrefixer("YourFunctionName")
	err := errors.New("the problem")
	fmt.Println(prefixer(err, "failed to whatever"))
	// Output: YourFunctionName: failed to whatever: the problem
}

func ExampleDeferErrorHandler() {
	_ = func() (err error) {
		file, err := os.Open("file.txt")
		if err != nil {
			return
		}
		defer DeferErrorHandler(file.Close(), &err)
		_, err = ioutil.ReadAll(file)
		return
	}()
}
