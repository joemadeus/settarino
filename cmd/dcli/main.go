package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nytimes/settarino/catalog"
	"github.com/nytimes/settarino/sets"
)

func main() {
	var loadedPSet *sets.PrimitiveSet
	var loadedFile string

	cmdread := bufio.NewReader(os.Stdin)
	cmdargs := make([]string, 0)
	cmd := ""

	for {
		switch cmd {
		case "load":
			if len(cmdargs) == 0 {
				fmt.Println("a file name is required")
				break
			}
			var err error
			loadedPSet, err = loadpset(cmdargs[0])
			if err != nil {
				fmt.Println(err)
				break
			}
			loadedFile = cmdargs[0]
			fmt.Printf("Set with tag %s was loaded from file %s\n", *loadedPSet.Tag(), loadedFile)
		case "loaded":
			fmt.Printf("Currently loaded set has tag %s, file %s\n", *loadedPSet.Tag(), loadedFile)
		case "unload":
			loadedPSet = nil
			loadedFile = ""
			fmt.Println("done")

		case "keys":
			if loadedPSet == nil {
				fmt.Println("please <load> a pset first")
				break
			}
			c := 25
			if len(cmdargs) > 0 {
				var err error
				c, err = strconv.Atoi(cmdargs[0])
				if err != nil {
					fmt.Println("the argument to 'keys' should be numeric")
					break
				}
			}
			i := 0
			for e := range loadedPSet.Elements() {
				fmt.Println(e.Key)
				i++
				if i > c {
					break
				}
			}
		case "stats":
			if loadedPSet == nil {
				fmt.Println("please load a pset first")
				break
			}
			fmt.Printf("  pset: %s\n  keys: %d\n  last update: %s\n",
				*loadedPSet.Tag(), loadedPSet.Size(), loadedPSet.LastUpdateTime().Format(time.RFC3339))

		case "exit":
			os.Exit(0)
		case "quit":
			os.Exit(0)
		default:
			fmt.Println(`dcli: Load a Primitive Set (pset) from a source and display information about it
Usage:
    load <filename>: load a pset from the specified file
    loaded: emit the name of the file that houses the currently loaded pset
    unload: remove the currently loaded pset from memory
    keys <count>: emit the first <count> keys from the loaded pset (defaults to 25)
    stats: emit statistics for the currently loaded pset

    quit | exit: byeeeeeeee`)
		}

		cmd = ""
		cmdargs = []string{}

		fmt.Print("> ")
		in, err := cmdread.ReadString('\n')
		if err != nil {
			_, _ = fmt.Fprint(os.Stderr, err.Error())
			os.Exit(-1)
		}

		in = strings.TrimSuffix(in, "\n")

		if len(in) == 0 {
			continue
		}

		insplit := strings.Split(in, " ")
		cmd = insplit[0]
		if len(insplit) > 1 {
			cmdargs = insplit[1:]
		}
	}
}

func loadpset(name string) (*sets.PrimitiveSet, error) {
	rd, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rd.Close(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, err.Error())
		}
	}()

	rb := bufio.NewReader(rd)
	return catalog.LoadSet(rb)
}
