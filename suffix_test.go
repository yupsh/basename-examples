package basename_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/basename"
)

func ExampleBasename_suffix() {
	// basename /path/to/file.txt .txt
	yup.MustRun(
		Basename("/path/to/file.txt", ".txt"),
	)
	// Output:
	// file
}

