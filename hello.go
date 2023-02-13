package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func getFormattedUrl() (string, error) {
	var inputReader *bufio.Reader
	var input string
	var formattedInput string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("What Website do you want to scan?\n")
	input, err = inputReader.ReadString('\n')

	if err == nil {
		formattedInput = strings.ReplaceAll(input, "\n", "/robots.txt")
		return "https://" + formattedInput, nil
	}
	return "", err
}
func main() {
	url, err := getFormattedUrl()
	if err != nil {
		fmt.Printf("Fatal Error in formatting that url properly: %s", err)
		return
	}
	robotsFile, err := http.Get(url)
	if err != nil {
		fmt.Printf("Fatal Error in Retrieving that URL's Robots.txt file: %s", err)
	}
	fmt.Printf("Robots file: %s", robotsFile.Body)

}
