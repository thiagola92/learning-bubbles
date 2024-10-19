package main

import (
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
)

type LaStopwatch struct {
	stopwatch stopwatch.Model
}

func newStopwatch() LaStopwatch {
	return LaStopwatch{
		stopwatch: stopwatch.New(),
	}
}

func (m LaStopwatch) Init() tea.Cmd {
	return m.stopwatch.Init()
}

func (m LaStopwatch) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	m.stopwatch, cmd = m.stopwatch.Update(msg)
	return m, cmd
}

func (m LaStopwatch) View() string {
	return m.stopwatch.View()
}
