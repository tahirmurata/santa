// Package elf provides helpers for solving the Advent of Code
package elf

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Elf struct {
	token string
}

// New creates a new Elf with the given token
func New(token string) *Elf {
	return &Elf{token: token}
}

// GetInputFile gets the input file for the given year and day
func (e *Elf) GetInputFile(year int, day int) *os.File {
	filename := fmt.Sprintf("%d-%d.txt", year, day)
	file, err := os.Open(filename)

	if errors.Is(err, os.ErrNotExist) {
		// handle the case where the file doesn't exist

		// Create file
		file_created, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file_created.Close()

		// Prepare request
		url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			panic(err)
		}
		request.AddCookie(&http.Cookie{Name: "session", Value: e.token})

		// Send request
		resp, err := http.DefaultClient.Do(request)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Write response to file
		_, err = io.Copy(file_created, resp.Body)
		if err != nil {
			panic(err)
		}

		// Open file
		file, err = os.Open(filename)
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	}

	return file
}