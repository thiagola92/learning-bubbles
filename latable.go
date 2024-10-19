package main

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type LaTable struct {
	table table.Model
}

func newTable() LaTable {
	t := table.New(
		table.WithColumns([]table.Column{
			{Title: "Name", Width: 16},
			{Title: "Age", Width: 3},
			{Title: "Tester", Width: 5},
		}),
		table.WithRows([]table.Row{
			{"Thiago", "31", "true"},
			{"Somebody", "18", "false"},
			{"No one", "200", "false"},
			{"More", "20", "false"},
			{"Two", "60", "false"},
		}),
		table.WithHeight(5),
	)
	t.Focus()

	return LaTable{
		table: t,
	}
}

func (m LaTable) Init() tea.Cmd {
	return nil
}

func (m LaTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		m.table, cmd = m.table.Update(msg)
	}

	return m, cmd
}

func (m LaTable) View() string {
	return m.table.View()
}
