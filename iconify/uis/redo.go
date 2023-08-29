package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 11c-.6 0-1 .4-1 1c0 2.9-1.5 5.5-4 6.9c-3.8 2.2-8.7.9-10.9-2.9C2.9 12.2 4.2 7.3 8 5.1c3.3-1.9 7.3-1.2 9.8 1.4h-2.4c-.6 0-1 .4-1 1s.4 1 1 1h4.5c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1s-1 .4-1 1v1.8C17 3 14.6 2 12 2C6.5 2 2 6.5 2 12s4.5 10 10 10s10-4.5 10-10c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}