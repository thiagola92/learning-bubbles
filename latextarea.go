package main

import (
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type LaTextArea struct {
	textarea textarea.Model
}

func newTextArea() LaTextArea {
	ta := textarea.New()
	ta.Focus()

	return LaTextArea{
		textarea: ta,
	}
}

func (m LaTextArea) Init() tea.Cmd {
	return textarea.Blink // Optional
	// return nil
}

func (m LaTextArea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		m.textarea, cmd = m.textarea.Update(msg)
	case cursor.BlinkMsg:
		m.textarea, cmd = m.textarea.Update(msg)
	}

	return m, cmd
}

func (m LaTextArea) View() string {
	return m.textarea.View()
}
