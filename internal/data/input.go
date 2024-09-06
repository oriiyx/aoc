package data

import (
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
		Value: "53616c7465645f5f2ea8c6960a8139b9b6c5f9ce68a06e7ceef2420e0fe5c264b61a124d3c06c7fc49a5c25a53b91abb60beb97e51a04873ab0a35c0bcaa1173",
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
