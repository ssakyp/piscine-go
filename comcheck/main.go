package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	// Flag to track if a match has been found
	found := false

	for _, e := range args {
		switch e {
		case "01", "galaxy", "galaxy 01":
			// Set found flag to true if a match is found
			found = true
		}
	}

	// If at least one match is found, print "Alert!!!"
	if found {
		os.Stdout.Write([]byte("Alert!!!\n"))
	}
}
