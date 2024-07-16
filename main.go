package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kingpin/v2"
)

var (
	filename = kingpin.Arg("filename", "Fortune filename").Default(".").String()
)

// Read filename from command line, and check if this filename is file, after
// that  just printFortune. If filenmae is directory, we get all filenames from
// directory and check if this file is fortune we adding filename to slice.
// After that we get random filename and then printForune with this filename.
func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	const fortunesDir = "fortunes"
	fi, err := os.Stat(fortunesDir)
	if err != nil {
		log.Fatal(fortunesDir + " is not exist...")
	}
	// if fi.Mode().IsRegular() {
	// 	if isFortune(*filename) {
	// 		printFortune(*filename)
	// 	}
	// } else
	if fi.Mode().IsDir() {
		files, err := os.ReadDir(fortunesDir)
		if err != nil {
			log.Fatal(err)
		}

		var forts []string
		for _, file := range files {
			filePath := fortunesDir + "/" + file.Name()
			if isFortune(filePath) {
				forts = append(forts, filePath)
			}
		}
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		rnd := r1.Intn(len(forts))
		printFortune(forts[rnd])
	}
}

// First version printFortune.
// Just read full file to memory, after that split to '/n%/n' and
// printing random fortunes.
//
// TODO We need to change this to reading a part of file.
func printFortune(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Cannot open " + filename + " filename...")
	}

	fortunes := strings.Split(string(data), "\n%\n")
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rnd := r1.Intn(len(fortunes))

	fmt.Println(fortunes[rnd])
}

// If we have filename and file with filename.dat, we then we think that this
// fortune file.
// We check that the filename exists, and filename.dat exist too
func isFortune(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	if _, err := os.Stat(filename + ".dat"); os.IsNotExist(err) {
		return false
	}
	return true
}
