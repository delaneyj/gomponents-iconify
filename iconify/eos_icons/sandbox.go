package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<circle cx="9.54" cy="8.99" r="1" fill="currentColor"/><circle cx="11.04" cy="11.99" r="1" fill="currentColor"/><circle cx="8.04" cy="11.99" r="1" fill="currentColor"/><path fill="currentColor" d="M19 9.24V15h-4a3 3 0 0 1-6 0H5V5h9.77l2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.24Z"/><path fill="currentColor" d="m17 10.1l-.84-.84l-.26.26l-1.44-1.41l.26-.26l-.78-.78c-.37-.38-.88-.17-1.15.31l-1.17 1.91a2 2 0 0 1 1.49 1.16a1.54 1.54 0 0 1 .6-.13A1.67 1.67 0 0 1 15.39 12l1.25-.74c.5-.26.71-.78.36-1.16Z"/><path fill="currentColor" d="m15.87 9.53l.26-.27l5.28-5.27l-1.42-1.42l-5.27 5.28l-.26.26l1.41 1.42z"/>`),
		g.Group(children),
	)
}