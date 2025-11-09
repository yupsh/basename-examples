package basename_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/basename"
)

func ExampleBasename_suffix() {
	// basename /path/to/file.txt .txt
	gloo.MustRun(
		Basename("/path/to/file.txt", ".txt"),
	)
	// Output:
	// file
}
