package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/joliv/spark"
)

func parse(line string) (float64, error) {
	split := strings.IndexFunc(line, unicode.IsSpace)
	if split == -1 {
		// no whitespace, interpret whole line as number
		split = len(line)
	}
	s := line[:split]

	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func render(data []float64) string {
	if len(data) < 2 {
		// unique little snowflake, no relative information
		return "❄"
	}
	return spark.Line(data)
}

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage:\n")
	fmt.Fprintf(os.Stderr, "  %s [NUMBER..]\n", prog)
	fmt.Fprintf(os.Stderr, "  %s <FILE\n", prog)
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()

	var data []float64
	if flag.NArg() == 0 {
		// read stdin, interpret first word of each line
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				continue
			}
			num, err := parse(line)
			if err != nil {
				log.Fatalf("cannot parse number: %s", err)
			}
			data = append(data, num)
		}
		if err := scanner.Err(); err != nil {
			log.Fatalf("cannot read standard input: %s", err)
		}
		if len(data) == 0 {
			log.Fatal("no data in stdin")
		}
	} else {
		for _, line := range flag.Args() {
			num, err := parse(line + "\n")
			if err != nil {
				log.Fatalf("cannot parse number: %s", err)
			}
			data = append(data, num)
		}
	}

	_, err := fmt.Println(render(data))
	if err != nil {
		log.Fatalf("cannot write to standard output: %s", err)
	}
}
