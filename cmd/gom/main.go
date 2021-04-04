package main

import (
	"flag"
	"fmt"
	"internal/metadata"
)

func main() {
	var helpFlag bool
	var helpFlagLong bool

	flag.BoolVar(&helpFlag, "h", false, "Get help")
	flag.BoolVar(&helpFlagLong, "help", false, "Get help")

	flag.Parse()
	args := flag.Args()
	if helpFlagLong || helpFlag {
		help()
	}
	fmt.Println("Arguments provided to goMetadata command: ", args)
	for i, arg := range args {
		fmt.Println("File ", i, ": ", arg)
		fmt.Println(metadata.GetMetadata(arg))
	}
}

func help() {
	helpMessage := "GoMetadata Tool Help\n" +
		"GoMetadata can be used to easily to get more accurate id3 tag metadata for your audio files. id3 metadata is taken from the itunes API.\n\n" +
		"SYNTAX:\n" +
		"COMMAND [options] [path to audio file]\n\n" +
		"OPTIONS:\n" +
		"-h/-help show help for the command usage."
	fmt.Println(helpMessage)
}

func getFlags(args []string) []string {
	var flags []string
	for _, arg := range args {
		var firstChar string = arg[0:1]
		if firstChar == "-" {
			flags = append(flags, arg)
		} else {
			break
		}
	}
	return flags
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
