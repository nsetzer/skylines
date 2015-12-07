
// Read file listing buildings one per line and output Critical Points
package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"log"
	"github.com/nsetzer/skylines/skylines"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage \n%s input_file\n" +
		"\tinput_file : file path or - for stdin\n" +
		"\tfile format is 3 integers per line\n" +
		"\tleft, right, height of each building.\n\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.Usage = Usage
}

func readFile(rd io.Reader) (buildings skylines.Buildings){
	scanner := bufio.NewScanner(rd)

	var l,r,h int

	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(),"%d %d %d",&l, &r, &h)
		buildings = append(buildings,skylines.Building{l,r,h})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {

	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Missing Positional Argument: input file:\n",
			os.Args[0])
		flag.Usage();
		os.Exit(1)
	}

	file_path := flag.Args()[0]

	var file *os.File
	var err error

	if file_path == "-" {
		file = os.Stdin
	} else {
		file,err = os.Open(file_path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	buildings := readFile(file)

	points := skylines.SolveFast(buildings)

	for _,p := range points {
		fmt.Println(p)
	}

}