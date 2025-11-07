package basename_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/basename"
)

func ExampleBasename_basic() {
	// basename /path/to/file.txt
	yup.MustRun(
		Basename("/path/to/file.txt"),
	)
	// Output:
	// file.txt
}

