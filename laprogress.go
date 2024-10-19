package main

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type LaProgress struct {
	progress progress.Model
}

func newProgress() LaProgress {
	p := progress.New()

	return LaProgress{
		progress: p,
	}
}

func (m LaProgress) Init() tea.Cmd {
	return nil
}

func (m LaProgress) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var newModel tea.Model

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		if msg.String() == "enter" {
			cmd = m.progress.SetPercent(m.progress.Percent() + 0.1)
		}
	case progress.FrameMsg:
		newModel, cmd = m.progress.Update(msg)
		m.progress = newModel.(progress.Model)
	}

	return m, cmd
}

func (m LaProgress) View() string {
	return m.progress.View()
}
