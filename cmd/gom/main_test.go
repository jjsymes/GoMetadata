package main

import (
	"testing"
)

func TestCanTest(t *testing.T) {
	var result int = 1
	if result != 1 {
		t.Error("Failed to test.")
	}
}

func ExampleMainWithNoOptionsOrArguments() {
	main()
	// Output:
	// ******* GOM *******
	// No files provided for processing, please provide a path to the command to use this tool.
}

func ExampleHelp() {
	help()
	// Output:
	// GoMetadata Tool Help
	// GoMetadata can be used to easily to get more accurate id3 tag metadata for your audio files. id3 metadata is taken from the itunes API.
	//
	// SYNTAX:
	// COMMAND [options] [path to audio file]
	//
	// OPTIONS:
	// -h/-help show help for the command usage.
}