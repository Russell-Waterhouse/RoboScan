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
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("What Website do you want to scan?\n")
	input, err := inputReader.ReadString('\n')
	handleFatalError(err, "ERROR READING INPUT")
	formattedInput := strings.ReplaceAll(input, "\n", "")
	return "https://" + formattedInput, nil
}

func main() {
	url, err := getFormattedUrl()
	handleFatalError(err, "ERROR FORMATTING THAT URL")
	response, err := http.Get(url + "/robots.txt")
	handleFatalError(err, "ERROR GETTING ROBOTS.TXT from URL: "+url)
	body, err := io.ReadAll(response.Body)
	handleFatalError(err, "ERROR READING BODY OF RESPONSE")
	err = response.Body.Close()
	handleFatalError(err, "ERROR CLOSING REQUEST CONNECTION")
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "Disallow") {
			directory := strings.ReplaceAll(scanner.Text(), "Disallow: ", "")
			requestUrl := url + directory
			response, err = http.Get(requestUrl)
			if err != nil {
				fmt.Printf("%s: Error: %s\n", requestUrl, err.Error())
				continue
			}
			fmt.Printf("%s: %s\n", requestUrl, response.Status)
		}
	}
}
