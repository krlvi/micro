package main

import (
	"github.com/gdamore/tcell"
	"strings"
)

const helpTxt = `Press Ctrl-q to quit help

Micro keybindings:

Ctrl-q:   Quit
Ctrl-s:   Save
Ctrl-o:   Open file

Ctrl-z:   Undo
Ctrl-y:   Redo

Ctrl-a:   Select all

Ctrl-c:   Copy
Ctrl-x:   Cut
Ctrl-v:   Paste

Ctrl-h:   Open help

Ctrl-u:   Half page up
Ctrl-d:   Half page down
PageUp:   Page up
PageDown: Page down

Ctrl-e:   Set option

Micro options:

colorscheme: loads the colorscheme stored in ~/.micro/colorschemes/'option'.micro
	default value: 'default'

tabsize: sets the tab size to 'option'
	default value: '4'
`

// DisplayHelp displays the help txt
// It blocks the main loop
func DisplayHelp() {
	topline := 0
	_, height := screen.Size()
	screen.HideCursor()
	totalLines := strings.Split(helpTxt, "\n")
	for {
		screen.Clear()

		lineEnd := topline + height
		if lineEnd > len(totalLines) {
			lineEnd = len(totalLines)
		}
		lines := totalLines[topline:lineEnd]
		for y, line := range lines {
			for x, ch := range line {
				st := tcell.StyleDefault
				screen.SetContent(x, y, ch, nil, st)
			}
		}

		screen.Show()

		event := screen.PollEvent()
		switch e := event.(type) {
		case *tcell.EventResize:
			_, height = e.Size()
		case *tcell.EventKey:
			switch e.Key() {
			case tcell.KeyUp:
				if topline > 0 {
					topline--
				}
			case tcell.KeyDown:
				if topline < len(totalLines)-height {
					topline++
				}
			case tcell.KeyCtrlQ, tcell.KeyCtrlW, tcell.KeyEscape, tcell.KeyCtrlC:
				return
			}
		}
	}
}