package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// This keystyle is For Color sakalaka boom boom
var (
	keystyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#93a2a3"))
)

type model struct {
	screen   bool
	quitting bool
}

type (
	frameMsg struct{}
)

func frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(time.Time) tea.Msg {
		return frameMsg{}
	})
}

func changeViewToMainScreen(m model) (tea.Model, tea.Cmd) {
	m.screen = true
	return m, frame()
}

func changeViewToAboutScreen(m model) (tea.Model, tea.Cmd) {
	m.screen = false
	return m, frame()
}

func mainView(m model) string {
	return fmt.Sprintf("\n \n Hi!! I AM %s \n\n\n", keystyle.Render(" ASHISH!"))

}

func aboutView(m model) string {
	return fmt.Sprintf("\n \n Hi!! I AM %s I mainly work on Backend %s\n\n\n", keystyle.Render(" Software Developer  "))

}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := message.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
		if k == "h" {
			return changeViewToAboutScreen(m)
		}
		if k == "l" {
			return changeViewToMainScreen(m)
		}

	}
	return m, nil
}

func (m model) View() string {
	var s string

	if !m.screen {
		s = mainView(m)
	} else {
		s = aboutView(m)
	}
	return s

}

func main() {

	p := tea.NewProgram(model{false, false}, tea.WithAltScreen())

	_, err := p.Run()

	if err != nil {
		log.Fatal(err)
	}
}
