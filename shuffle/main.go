package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var mode = flag.Int("mode", 0, "set != 0 to remove prefix")

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		panic("should be dir name")
	}

	dir := args[0]

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for _, file := range files {
		if file.IsDir() == false {
			name := file.Name()

			if *mode == 0 {
				// basic shuffle
				os.Rename(
					filepath.Join(dir, name),
					filepath.Join(dir, fmt.Sprintf("%d %s", r.Intn(1000), name)),
				)
			} else {
				words := strings.Split(name, " ")

				if len(words) == 1 {
					continue
				}

				// remove shuffle
				os.Rename(
					filepath.Join(dir, name),
					filepath.Join(dir, strings.Join(words[1:], " ")),
				)
			}
		}
	}
}
