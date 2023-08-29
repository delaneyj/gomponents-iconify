package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 5c-1.11 0-2 .89-2 2v10c0 1.11.89 2 2 2h16c1.11 0 2-.89 2-2V7c0-1.11-.89-2-2-2H4m.5 2a1 1 0 0 1 1 1a1 1 0 0 1-1 1a1 1 0 0 1-1-1a1 1 0 0 1 1-1M7 7h13v10H7V7m1 1v8h3V8H8m4 0v8h3V8h-3m4 0v8h3V8h-3M9 9h1v1H9V9m4 0h1v1h-1V9m4 0h1v1h-1V9Z"/>`),
		g.Group(children),
	)
}