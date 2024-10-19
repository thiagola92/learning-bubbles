package main

import tea "github.com/charmbracelet/bubbletea"

func main() {
	p := tea.NewProgram(
		newHelp(),
		tea.WithAltScreen(),       // Fullscreen mode
		tea.WithMouseCellMotion(), // Enable mouse motion
	)

	_, err := p.Run()

	if err != nil {
		print(err.Error())
	}
}
