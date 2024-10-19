package main

import (
	"github.com/charmbracelet/bubbles/filepicker"
	tea "github.com/charmbracelet/bubbletea"
)

type LaFilepicker struct {
	filepicker filepicker.Model
	path       string
}

func newFilepicker() LaFilepicker {
	fp := filepicker.New()

	return LaFilepicker{
		filepicker: fp,
	}
}

func (m LaFilepicker) Init() tea.Cmd {
	// Needs to be done here
	return m.filepicker.Init()
}

func (m LaFilepicker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}

	// Make filepicker process keys BEFORE checking if a file is selected
	m.filepicker, cmd = m.filepicker.Update(msg)
	selected, path := m.filepicker.DidSelectFile(msg)

	if selected {
		m.path = path
	}

	return m, cmd
}

func (m LaFilepicker) View() string {
	return m.filepicker.View() +
		"\n\nSelected file: " +
		m.path
}
