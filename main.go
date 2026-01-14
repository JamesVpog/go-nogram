package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type pos struct {
	col int 
	row int
}
type model struct {
	board [][]int
	solution [][]int
	cursor pos
	
	height int
	width int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// what messages were sent to update?	
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		
	case tea.KeyMsg:
		switch msg.String(){			
		// what was the keypress?
		case "ctrl+c", "q":
			return m, tea.Quit
		case "k", "up":
			// move the cursor up to the next row if possible
			if m.cursor.row > 0 {	
				m.cursor.row--
			}
		case "l", "right":
			// move right to next col if possible
			if m.cursor.col < len(m.board[0]) - 1 {
				m.cursor.col++
			}
		case "j", "down":
			if m.cursor.row < len(m.board) - 1 {
				m.cursor.row++
			}
		case "h", "left":
			if m.cursor.col > 0 {
				m.cursor.col--
			}
		case " ", "enter":
			// fill only if the solution allows to be filled
			if m.solution[m.cursor.row][m.cursor.col] == 1 {	
				// user was right!
				m.board[m.cursor.row][m.cursor.col] = 1	 
			} else {
				// user got it wrong!
				m.board[m.cursor.row][m.cursor.col] = 3
			}
		case "x", "X":
			// draw x in the square at that position 
			m.board[m.cursor.row][m.cursor.col] = 2
		}
	}
	return m, nil
}

// styles for the view
var (	
	baseBlack 	= lipgloss.Color("0")
	
	cyan	= lipgloss.Color("6")
	red    = lipgloss.Color("1")
	green  = lipgloss.Color("2")
	purple = lipgloss.Color("5")
	
	cellWidth = 10
	cellHeight = 5
	
	baseStyle = lipgloss.NewStyle().
	    Width(cellWidth).
	    Height(cellHeight).
	    Align(lipgloss.Center, lipgloss.Center).
	    Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("7"))
)


func (m model) View() string {
	
	renderCell := func(r, c int) string {
		isFocused := (r == m.cursor.row && c == m.cursor.col)
		currentState := m.board[r][c]

		style := baseStyle
		// 1. check if focused
		if isFocused {
			style = style.BorderForeground(cyan).Border(lipgloss.ThickBorder())
		}
		
		// switch on the data to return the appropriate 4 states of each nonogram cell
		// 0. empty, 
		// 1. filled, 
		// 2. user-x'd out, 
		// 3. incorrectly filled
		
		switch currentState {
		case 1:
			//filled in
			style = style.Foreground(green).Background(green)
		case 2:
			// marked as x by the user
			return style.Render("x")
		case 3:
			// fileld in but wrong!
			style = style.Foreground(red).Background(red)
		}
		return style.Render("")
	}
	// loop through the m.board and render each cell, store in renderedCells
	// use joinHorizontal to make a row, stored in rows
	var rows []string
	for r := range m.board {
		var renderedCells []string
		for c := range m.board[r] {
			cell := renderCell(r, c)
			renderedCells = append(renderedCells, cell)
		}
		row := lipgloss.JoinHorizontal(lipgloss.Top, renderedCells...)
		rows = append(rows, row)
	}
	
	// send rows to joinVertical to make the grid
	grid := lipgloss.JoinVertical(lipgloss.Left, rows...)
	
	//TODO: instructions on the bottom?
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, grid)
}

func initialModel() model {
	// example where filling in the middle square is the solution
	return model{
		board: [][]int{
			{0, 0, 0},
		 	{0, 0, 0},
			{0, 0, 0}},
		
		solution: [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0}},
	}
}

func main(){	
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
	
	
}
