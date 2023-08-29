package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExpandDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4V2H2v2h9v14.17l-5.5-5.5l-1.42 1.41L12 22l7.92-7.92l-1.42-1.41l-5.5 5.5V4h9Z"/>`),
		g.Group(children),
	)
}