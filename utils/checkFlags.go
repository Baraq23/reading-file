package utils

import (
	"errors"
	"flag"
)

// CheckFlags validates the flags passed to the program
func CheckFlags(args []string) error {
	fs := flag.NewFlagSet("textindex", flag.ContinueOnError)

	// Expected flags
	command := fs.String("c", "", "Specify the command (index or lookup)")
	input := fs.String("i", "", "Specify the input file")
	chunkSize := fs.String("s", "", "Specify the chunk size")
	output := fs.String("o", "", "Specify the output file")
	query := fs.String("q", "", "Specify the query text")

	if err := fs.Parse(args[1:]); err != nil {
		return errors.New("flag not used")
	}

	// Usage message for error handling
	const usageMsg = "usage: textindex  -c index  -i <input_file.txt>  -s <chunk-size>  -o <index_file.idx>\ntextindex  -c lookup  -i <index_file.idx> -q <query_text>"

	switch *command {
	case "index":
		if *input == "" || *chunkSize == "" || *output == "" {
			return errors.New(usageMsg)
		}
	case "lookup":
		if *input == "" || *query == "" {
			return errors.New(usageMsg)
		}
	default:
		return errors.New("flag not used")
	}

	return nil
}
