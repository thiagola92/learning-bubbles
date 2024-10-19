package main

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type LaList struct {
	list list.Model
}

type LaItem struct {
	field1 string
	field2 string
	field3 string
}

// Items from List must implement this method
func (i LaItem) FilterValue() string { return i.field1 }

// These must be implemented because "NewDefaultDelegate()" is being used
func (i LaItem) Title() string       { return i.field2 }
func (i LaItem) Description() string { return i.field3 }

func newList() LaList {
	l := list.New(
		[]list.Item{
			LaItem{field1: "hello", field2: "hello", field3: "meh1"},
			LaItem{field1: "world", field2: "world", field3: "meh2"},
			LaItem{field1: "how", field2: "how", field3: "meh3"},
			LaItem{field1: "are", field2: "are", field3: "meh4"},
			LaItem{field1: "you", field2: "you", field3: "meh5"},
			LaItem{field1: "?", field2: "?", field3: "meh6"},
			LaItem{field1: "I'm", field2: "I'm", field3: "meh7"},
			LaItem{field1: "fine", field2: "fine", field3: "meh8"},
			LaItem{field1: ",", field2: ",", field3: "meh9"},
			LaItem{field1: "thank", field2: "thank", field3: "meh10"},
			LaItem{field1: "you", field2: "you", field3: "meh11"},
		},
		list.NewDefaultDelegate(), 0, 0, // Size doesn't matter because we resize during Update()
	)
	l.Title = "LaList"

	return LaList{
		list: l,
	}
}

func (m LaList) Init() tea.Cmd {
	return nil
}

func (m LaList) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// Resize when window change (we are using full window size)
		m.list.SetSize(msg.Width, msg.Height)
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m LaList) View() string {
	return m.list.View()
}
