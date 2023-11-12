package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {

	// The command you want to run along with the argument
	cmd := exec.Command("git", "branch")

	// Get a pipe to read from standard out
	r, _ := cmd.StdoutPipe()

	// Use the same pipe for standard error
	cmd.Stderr = cmd.Stdout

	// Make a new channel which will be used to ensure we get all output
	done := make(chan struct{})

	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)

	// Use the scanner to scan the output line by line and log it
	// It's running in a goroutine so that it doesn't block
    var outputLines []string
	go func() {

		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
            fmt.Println(line)
            outputLines = append(outputLines, line)
		}

		// We're all done, unblock the channel
		done <- struct{}{}

	}()

	// Start the command and check for errors
	cmd.Start()

	// Wait for all output to be processed
	<-done

	// Wait for the command to finish
    fmt.Println(outputLines)
	cmd.Wait()
}

