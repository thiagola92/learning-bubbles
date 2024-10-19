package main

import (
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
)

type LaTimer struct {
	timer timer.Model
}

func newTimer() LaTimer {
	s, _ := time.ParseDuration("10s")
	t := timer.New(s)

	return LaTimer{
		timer: t,
	}
}

func (m LaTimer) Init() tea.Cmd {
	return m.timer.Init()
}

func (m LaTimer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case timer.TimeoutMsg:
		return m, tea.Quit
	case timer.TickMsg:
		m.timer, cmd = m.timer.Update(msg)
	}

	return m, cmd
}

func (m LaTimer) View() string {
	return m.timer.View()
}
