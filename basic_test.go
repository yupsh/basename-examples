package basename_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/basename"
)

func ExampleBasename_basic() {
	// basename /path/to/file.txt
	gloo.MustRun(
		Basename("/path/to/file.txt"),
	)
	// Output:
	// file.txt
}
