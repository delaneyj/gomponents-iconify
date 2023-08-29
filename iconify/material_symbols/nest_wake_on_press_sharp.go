package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWakeOnPressSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13V4h3v9h-3ZM7.1 21L2 15.625L3.2 14.4l3.8.85V4.5q0-.625.438-1.063T8.5 3q.625 0 1.063.438T10 4.5v6h1.4l5.775 2.9L16.1 21h-9Z"/>`),
		g.Group(children),
	)
}