package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func More(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13.5a1.5 1.5 0 0 1-1.5-1.5a1.5 1.5 0 0 1 1.5-1.5a1.5 1.5 0 0 1 1.5 1.5a1.5 1.5 0 0 1-1.5 1.5m-5 0a1.5 1.5 0 0 1-1.5-1.5a1.5 1.5 0 0 1 1.5-1.5a1.5 1.5 0 0 1 1.5 1.5a1.5 1.5 0 0 1-1.5 1.5m-5 0A1.5 1.5 0 0 1 7.5 12A1.5 1.5 0 0 1 9 10.5a1.5 1.5 0 0 1 1.5 1.5A1.5 1.5 0 0 1 9 13.5M22 3H7c-.69 0-1.23.35-1.59.88L0 12l5.41 8.11c.36.53.96.89 1.65.89H22a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}