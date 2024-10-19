package main

import (
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
)

type LaPaginator struct {
	paginator paginator.Model
}

func newPaginator() LaPaginator {
	p := paginator.New(
		paginator.WithTotalPages(7),
	)

	return LaPaginator{
		paginator: p,
	}
}

func (m LaPaginator) Init() tea.Cmd {
	return nil
}

func (m LaPaginator) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

		m.paginator, cmd = m.paginator.Update(msg)
	}

	return m, cmd
}

func (m LaPaginator) View() string {
	return m.paginator.View()
}
