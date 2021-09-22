package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ryantriangles/oshash"
)

var binFlag = flag.Bool("b", false, "display binary values")
var decFlag = flag.Bool("d", false, "display decimal values")
var hexFlag = flag.Bool("x", false, "display hexadecimal values")
var pipeFlag = flag.Bool("pipe", false, "pipe mode")
var showFilenames = flag.Bool("f", false, "show filenames with output")
var sep = flag.String("sep", "\t", "separator string for tabular output")

func init() {
	flag.Parse()

	// If no flags are used, set the hexadecimal flag to true, making
	// hexadecimal representation the default.
	if !*binFlag && !*decFlag {
		*hexFlag = true
	}
}

func handleFilename(filename string) {
	h, err := oshash.FromFilepath(filename)
	if err == oshash.ErrDataTooSmall {
		fmt.Println("Too small")
		return
	}
	if err != nil {
		panic(err)
	}
	str := []string{}
	if *showFilenames {
		str = append(str, filename)
	}
	if *hexFlag {
		str = append(str, strconv.FormatUint(h, 16))
	}
	if *binFlag {
		str = append(str, strconv.FormatUint(h, 2))
	}
	if *decFlag {
		str = append(str, strconv.FormatUint(h, 10))
	}
	fmt.Println(strings.Join(str, *sep))
}

func stdinScanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		filename := scanner.Text()
		handleFilename(filename)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func main() {
	if *pipeFlag {
		stdinScanner()
	} else {
		for _, arg := range flag.Args() {
			handleFilename(arg)
		}
	}
}
