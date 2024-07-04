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
	return lines
}

func GetInputString(requestURL string) string {
	cookie := &http.Cookie{
		Name:  "session",
		Value: "53616c7465645f5f0551105146976e526fb48d374cea918692d558dbb2789b0c5ba8dadcf3006ee793ea3ab270af21c8d6fe14c0d03a0b1ab5e87bcfff6bd1e6",
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
