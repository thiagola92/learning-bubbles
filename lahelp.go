package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// This will be used to test pressing UP
var counter = 0

// Bubbles help will need a struct that implement ShortHelp() and FullHelp()
type keyMap struct {
	Up   key.Binding
	Down key.Binding
	Help key.Binding
	Quit key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Help}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Quit}, // First column
		{k.Help},               // Second column
	}
}

var keys = keyMap{
	// Simplest example of creating a key binding
	// but without a WithHelp() it will not appear during
	// ShortHelp() and FullHelp()
	// Still occupies spaces in both cases...
	Up: key.NewBinding(
		key.WithKeys("up", "w"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "s"),
		key.WithHelp("↓/s", "decrement counter"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
}

type LaHelp struct {
	help help.Model
	keys keyMap
}

func newHelp() LaHelp {
	return LaHelp{
		help: help.New(),
		keys: keys,
	}
}

func (m LaHelp) Init() tea.Cmd {
	return nil
}

func (m LaHelp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		// This will make it show "..." when more text is available but not visible
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, m.keys.Up):
			counter += 1
			return m, nil
		case key.Matches(msg, m.keys.Down):
			counter -= 1
			return m, nil
		}
	}

	// m.help, cmd = m.help.Update(msg)
	return m, cmd
}

func (m LaHelp) View() string {
	return fmt.Sprintf("%s\nCounter: %d", m.help.View(m.keys), counter)
}
