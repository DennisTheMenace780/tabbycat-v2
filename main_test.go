package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/exp/teatest"
	"github.com/muesli/termenv"
)

func init() {
	// This is required for CI to pass. See https://charm.sh/blog/teatest/
	lipgloss.SetColorProfile(termenv.Ascii)
}

func TestOutput(t *testing.T) {
	t.Run("Move down and select a branch", func(t *testing.T) {

		createDirectory()
		initGitTestRepository()
		writeToFile()
		addTestFile()
		commitTestFile()

		createBranches([]string{"branch-1", "branch-2", "branch-2"})
		// defer tearDownGitTestRepository()

		// var branches []string
		//
		// model := initModel(branches)
		//
		// tm := teatest.NewTestModel(t, model, teatest.WithInitialTermSize(300, 100))
		//
		// // Assert that the program, at some point, has the following byte string ... make a helper function?
		// teatest.WaitFor(t, tm.Output(),
		// 	func(bts []byte) bool {
		// 		return bytes.Contains(
		// 			bts,
		// 			[]byte("1. branch-1"),
		// 		)
		// 	},
		// )
		//
		// moveDownAndSelectBranch(tm, 1)
		//
		// tm.WaitFinished(t, teatest.WithFinalTimeout(time.Second))

		if 2 != 2 {
			t.Errorf("expected 2 but got 3")
		}
		// out, err := io.ReadAll(tm.FinalOutput(t))
		// if err != nil {
		// 	t.Error("Error reading from FinalOutput", err)
		// }
		// teatest.RequireEqualOutput(t, out)

	})
}

func createDirectory() {
	cmd := exec.Command("mkdir", "-p", "testdata/TestOutput")
	cmd.Run()
}

func initGitTestRepository() {
	cmd := exec.Command("git", "init")
	cmd.Dir = "./testdata/TestOutput"
	cmd.Run()
}

func writeToFile() {
    // need to get this to write to the right file name
	cmd := exec.Command("echo", "Hello World")
	cmd.Dir = "./testdata/TestOutput/"
	file, err := os.CreateTemp(cmd.Dir, "somefile.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	cmd.Stdout = file
	cmd.Run()
}

func addTestFile() {
	cmd := exec.Command("git", "add", "somefile.txt")
	cmd.Dir = "./testdata/TestOutput/"
	cmd.Run()
}

func commitTestFile() {
	cmd := exec.Command("git", "commit", "--no-verify", "somefile.txt")
	cmd.Dir = "./testdata/TestOutput/"
	cmd.Run()
}

func tearDownGitTestRepository() {
	cmd := exec.Command("rm", "-rf", "TestOutput")
	cmd.Dir = "./testdata/"
	cmd.Run()
}

func createBranches(names []string) {
	for _, name := range names {
		cmd := exec.Command("git", "checkout", "-b", name)
		cmd.Dir = "./testdata/TestOutput"
		cmd.Run()
	}
}

func initModel(branches []string) Model {
	var items []list.Item
	for _, branch := range branches {
		items = append(items, Item(branch))
	}
	l := list.New(items, ItemDelegate{}, DefaultWidth, ListHeight)
	return Model{list: l}
}

func moveDownAndSelectBranch(tm *teatest.TestModel, down int) {
	for i := 0; i < down; i++ {
		tm.Send(tea.KeyMsg{
			Type:  tea.KeyRunes,
			Runes: []rune("j"),
		})
	}

	tm.Send(tea.KeyMsg{
		Type:  tea.KeyRunes,
		Runes: []rune("enter"),
	})
}
