package structs

import (
	tea "github.com/charmbracelet/bubbletea"
)

type CommandModel struct {
	Choices        []string
	Cursor         int
	Selected       map[int]struct{}
	ViewCallback   func() string
	UpdateCallback func(tea.Msg) (tea.Model, tea.Cmd)
}

type Command struct {
	Name         string
	Type         string
	DefaultValue string
	Flag         string
	Weight       int
	Description  string
}

func (c Command) ParseDefault() interface{} {
	switch c.Type {
	case "bool":
		return c.DefaultValue == "true"
	case "string":
		return c.DefaultValue
	default:
		return nil
	}
}



func (m CommandModel) Init() tea.Cmd {
	return nil
}

func (m CommandModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.UpdateCallback != nil {
		return m.UpdateCallback(msg)
	}
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m CommandModel) View() string {
	if m.ViewCallback != nil {
		return m.ViewCallback()
	}
	return ""
}
