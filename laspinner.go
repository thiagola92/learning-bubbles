package main

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type LaSpinner struct {
	spinner spinner.Model
}

func newSpinner() LaSpinner {
	return LaSpinner{
		spinner: spinner.New(),
	}
}

func (m LaSpinner) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m LaSpinner) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	default:
		m.spinner, cmd = m.spinner.Update(msg)
	}

	return m, cmd
}

func (m LaSpinner) View() string {
	return m.spinner.View()
}
