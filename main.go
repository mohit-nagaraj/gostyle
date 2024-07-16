package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/mohit-nagaraj/gowow/catlol"
	"github.com/mohit-nagaraj/gowow/cowsay"
)

var (
	filename = kingpin.Arg("filename", "Fortune filename").Default(".").String()
)

// main is the entry point of the application. It sets up the command line arguments,
// verifies the existence of the fortunes directory, selects a random fortune file,
// retrieves a random fortune from the file, and then formats and displays it using
// a cow image with colored text.
func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	const fortunesDir = "fortunes"
	fi, err := os.Stat(fortunesDir)
	if err != nil {
		log.Fatal(fortunesDir + " does not exist...")
	}
	// Check if the fortunesDir is a directory.
	if fi.Mode().IsDir() {
		files, err := os.ReadDir(fortunesDir) // Read the contents of the directory.
		if err != nil {
			log.Fatal(err)
		}

		var forts []string
		for _, file := range files {
			filePath := fortunesDir + "/" + file.Name() // Check if the file is a fortune file.
			if isFortune(filePath) {
				forts = append(forts, filePath)
			}
		}
		s1 := rand.NewSource(time.Now().UnixNano()) // Select a random fortune file.
		r1 := rand.New(s1)                          // Retrieve a random fortune from the file.
		rnd := r1.Intn(len(forts))
		fortune := getFortune(forts[rnd])
		balloon := cowsay.BuildBalloon(fortune)                 //display cow
		rainbow := catlol.MakeRainbowText(balloon + cowsay.Cow) //color the text in rainbow colors.
		fmt.Println(rainbow)
	}
}

// Reads the content of a fortune file, splits it into individual fortunes,
// and returns a random fortune.
//
// filename: The path to the fortune file.
func getFortune(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal("Cannot open " + filename + " filename...")
	}
	// Split the content by "\n%\n" to separate individual fortunes.
	fortunes := strings.Split(string(data), "\n%\n")
	// Select a random fortune.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rnd := r1.Intn(len(fortunes))

	return fortunes[rnd]
}

// Checks if a given filename is a valid fortune file. A valid fortune file
// has a corresponding ".dat" file.
//
// filename: The name of the file to check.
func isFortune(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}

	if _, err := os.Stat(filename + ".dat"); os.IsNotExist(err) {
		return false
	}
	return true
}
