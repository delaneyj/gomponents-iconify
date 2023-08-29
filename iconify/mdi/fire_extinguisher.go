package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireExtinguisher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 7h1.25L12 5h-1.75L6 7.5V9H4V6.5L10 3h2V2h2v1h2l1-.5v3L16 5h-2l.25 2h1.25A1.5 1.5 0 0 1 17 8.5V22H9V8.5A1.5 1.5 0 0 1 10.5 7Z"/>`),
		g.Group(children),
	)
}