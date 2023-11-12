package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	list     list.Model
	choice   string
	err      string
	quitting bool
}

func (m Model) Init() tea.Cmd {
	return nil
}

func processBranchString(s string) string {
	str := strings.TrimLeft(s, "*")
	return strings.TrimSpace(str)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = processBranchString(string(i))

				out, err := exec.Command("git", "checkout", "some-branch-1").CombinedOutput()
				if err != nil {
					log.Print("TEST", err)
				}
				log.Printf("Output: \n%s", out)
                m.err = string(out)
			}
			return m, tea.Quit
		}

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() string {
    if m.err != "" {
        return QuitCheckoutStyle.Render(m.err)
    }
	if m.choice != "" {
		return QuitTextStyle.Render(fmt.Sprintf("switched to branch '%s'", m.choice))
	}
	if m.quitting {
		return QuitTextStyle.Render("Not hungry? That’s cool.")
	}
	return "\n" + m.list.View()
}
