package parser

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getFromCache(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return file
}

func saveToCache(path string, input []byte) error {

	if err := os.WriteFile(path, input, 0644); err != nil {
		return err
	}
	return nil
}
func getFromWeb(day int) []byte {
	uri := fmt.Sprintf("https://adventofcode.com/2022/day/%d/input", day)

	req, err := http.NewRequest(http.MethodGet, uri, nil)

	if err != nil {
		panic("error getting the puzzle input")
	}

	muhCookie := os.Getenv("AOC_COOKIE")

	if muhCookie == "" {
		panic("cookie not set")
	}

	aocCookie := http.Cookie{
		Name:  "session",
		Value: muhCookie,
	}
	req.AddCookie(&aocCookie)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	return body

}

func getInput(day int) string {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := fmt.Sprintf("%s/input/input_%d", dir, day)
	var input []byte

	if _, err := os.Stat(path); err != nil {
		input = getFromWeb(day)
		if err := saveToCache(path, input); err != nil {
			panic(err)
		}
	} else {
		input = getFromCache(path)
	}
	return strings.TrimSuffix(string(input), "\n")
}
func ReadInputFile(day int) string {
	return getInput(day)
}
