package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func handleFatalError(err error, errorMessage string) {
	if err == nil {
		return
	}
	fmt.Printf("Fatal Error: %s\n", errorMessage)
	fmt.Printf("Debug message: %s\n", err.Error())
	fmt.Printf("\n\nExiting the Program\n\n")
	os.Exit(-1)
}

func getFormattedUrl() (string, error) {
	var inputReader *bufio.Reader
	var input string
	var formattedInput string
	var err error

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Printf("What Website do you want to scan?\n")
	input, err = inputReader.ReadString('\n')
	handleFatalError(err, "ERROR READING INPUT")
	formattedInput = strings.ReplaceAll(input, "\n", "/robots.txt")
	return "https://" + formattedInput, nil
}
func main() {
	url, err := getFormattedUrl()
	handleFatalError(err, "ERROR FORMATTING THAT URL")
	response, err := http.Get(url)
	handleFatalError(err, "ERROR GETTING ROBOTS.TXT from URL: "+url)
	body, err := io.ReadAll(response.Body)
	handleFatalError(err, "ERROR READING BODY OF RESPONSE")
	fmt.Println(string(body))
	err = response.Body.Close()
	handleFatalError(err, "ERROR CLOSING REQUEST CONNECTION")
}
