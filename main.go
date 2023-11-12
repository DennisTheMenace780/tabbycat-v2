package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cmd := exec.Command("git", "branch")
	// Get a pipe to read from standard out
	r, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	// Make a new channel which will be used to ensure we get all output
	done := make(chan struct{})
	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)
	// Use the scanner to scan the output line by line and log it
	// It's running in a goroutine so that it doesn't block
	var branches []string
	go func() {
		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
			branches = append(branches, line)
		}
		// We're all done, unblock the channel
		done <- struct{}{}
	}()
	// Start the command and check for errors
	cmd.Start()

	// Wait for all output to be processed
	<-done
	// Wait for the command to finish
	fmt.Println(branches)

	branchItems := BuildItems(branches)
	l := ListBuilder(branchItems)

	if _, err := tea.NewProgram(Model{list: l}).Run(); err != nil {
		log.Fatal(err)
	}
	cmd.Wait()
}
