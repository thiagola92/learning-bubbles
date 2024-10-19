package main

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type LaViewport struct {
	viewport viewport.Model
}

func newViewport() LaViewport {
	v := viewport.New(10, 10)
	v.SetContent(
		"Hello, how are you? I'm fine, just testing this viewport. May it will let me scroll horizontally\nBut\nmaybe\nnot\n...\nMaybe\nVertically?" +
			"More\ntext\nis\nalways\nwelcome\nhere\n...\nbecause\nI\nstill\nneed\nto\ntest\nthe\nscrooling\n",
	)
	// v.HighPerformanceRendering = true

	return LaViewport{
		viewport: v,
	}
}

func (m LaViewport) Init() tea.Cmd {
	return nil
}

func (m LaViewport) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		// case tea.WindowSizeMsg:
		// 	cmds = append(cmds, viewport.Sync(m.viewport))
	}

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m LaViewport) View() string {
	return m.viewport.View()
}
