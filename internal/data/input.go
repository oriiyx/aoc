package data

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

func GetInputArray(requestURL string) []string {
	input := GetInputString(requestURL)
	lines := strings.Split(input, "\n")
	lines = deleteEmptyStrings(lines)

	return lines
}

func GetInputString(requestURL string) string {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5ffb7b0ff56b570baed4844245e1b6c6234c66f20561ea732fd6b4bde725d899fd85cba4a26d966f0d93d7031d3f4d8b56112f9fbd7d5c5817",
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Printf("error creating cookie jar: %s\n", err)
		os.Exit(1)
	}

	client := &http.Client{
		Jar:       jar,
		Transport: &http.Transport{},
	}

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	client.Jar.SetCookies(req.URL, []*http.Cookie{cookie})

	resp, responseErr := client.Do(req)
	if responseErr != nil {
		fmt.Printf("error getting response: %s\n", responseErr)
		os.Exit(1)
	}
	defer resp.Body.Close()

	resBody, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Printf("error reading http request: %s\n", readErr)
		os.Exit(1)
	}

	return string(resBody)
}

func deleteEmptyStrings(inputs []string) []string {
	var returnValues []string
	for _, str := range inputs {
		if str != "" {
			returnValues = append(returnValues, str)
		}
	}
	return returnValues
}

// ReadLines reads a file and returns its contents as a string slice,
// with each line as a separate element.
func ReadLines(filename string) ([]string, error) {
	dir, err := os.Getwd()

	file, err := os.Open(dir + "/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Trim any Windows-style line endings
		line := strings.TrimRight(scanner.Text(), "\r")
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
