/*
Bing Search Automation (Go)

This program automates Bing searches using terms from text files.
It can help simulate Microsoft Rewards points accumulation by opening searches in your browser.

Usage Examples:

# Perform 60 searches using a random bing-0.txt to bing-5.txt file get 180 MS POINTS (XMAS Promo)
go run main.go -num=50 -single=false -debug=true

# Perform 30 searches using the default bing.txt file to get 120 MS POINTS
go run main.go -num=30 -single=true -debug=false

Flags:
-num    (required) Number of searches to perform. Must be > 0
-single (optional) Use a random file (true) or default file (false). Default: true
-debug  (optional) Shorten sleep and avoid opening the broweser: Default: True

Notes:
- Each search term must be on a separate line in the text file.
- A 10-second delay is applied between searches to avoid anti-bot triggers.
- Ensure the .txt files exist in the same directory as the program.
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/pkg/browser"
)

type Rewards struct {
	num   int
	file  string
	url   string
	debug bool
}

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func cmdHelper() (Rewards, error) {
	num := flag.Int("num", 0, "number of searches (e.g. -num=120)")
	single := flag.Bool("single", true, "use a single random file or default file (e.g. -single=true)")
	debug := flag.Bool("debug", true, "use debug to showcase the outout (e.g. -debug=true)")

	flag.Parse()

	if *num <= 0 {
		return Rewards{}, fmt.Errorf("invalid -num value: must be greater than 0")
	}

	var fileName string

	if *single {
		// Use a default file
		fileName = "bing.txt"
	} else {
		// Randomly pick one of bing-0.txt to bing-5.txt
		fileName = "bing-" + strconv.Itoa(rand.Intn(6)) + ".txt"
	}

	rewards := Rewards{
		num:   *num,
		file:  fileName,
		url:   "https://bing.com/search?q=",
		debug: *debug,
	}

	return rewards, nil
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open file %q: %w", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan file %q: %w", path, err)
	}

	return words, nil
}

func main() {

	microsoft, errCMD := cmdHelper()

	if errCMD != nil {
		log.Fatalf("parse arguments: %v", errCMD)
	}

	words, errRF := readFile(microsoft.file)

	if errRF != nil {
		log.Fatalf("read file: %v", errRF)
	}

	fmt.Println(Cyan + "\n\n[X] MS Rewards - Bing Engine Golang Configuration [X]\n" + Reset)
	fmt.Println(Yellow+"[-] Search:", microsoft.num, Reset)
	fmt.Println(Yellow+"[-] Source:", microsoft.file, Reset)
	fmt.Println(Yellow+"[-] Computing MS Rewards Points:", 3*microsoft.num, Reset)
	fmt.Print("\n")
	fmt.Println(Green+"[-] Log search:\n", Reset)

	// Create a copy of the words slice indices and shuffled indices
	indices := rand.Perm(len(words))

	sleepDuration := 10 * time.Second

	if microsoft.debug {
		sleepDuration = 1 * time.Second
	}

	for i := 0; i < microsoft.num && i < len(words); i++ {
		idx := indices[i]
		query := url.QueryEscape(words[idx])
		bingURL := microsoft.url + query + "&PC=U316&FORM=CHROMN"

		fmt.Printf(" + [%d][%s]: %s\n", i, words[idx], bingURL)
		if !microsoft.debug {
			// Open browser only if not in debug mode
			err := browser.OpenURL(bingURL)
			if err != nil {
				fmt.Printf("%s[!] Failed to open browser: %v%s\n", Red, err, Reset)
			}
		}
		time.Sleep(sleepDuration)
	}
}
