package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type LaTextInput struct {
	textinput textinput.Model
}

func newTextinput() LaTextInput {
	ti := textinput.New()
	ti.Focus()

	return LaTextInput{
		textinput: ti,
	}
}

func (m LaTextInput) Init() tea.Cmd {
	return textinput.Blink // Optional
	// return nil
}

func (m LaTextInput) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		m.textinput, cmd = m.textinput.Update(msg)
	}

	return m, cmd
}

func (m LaTextInput) View() string {
	return m.textinput.View()
}
