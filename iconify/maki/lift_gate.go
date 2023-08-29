package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LiftGate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M1 4a2 2 0 1 1 4 0v8a1 1 0 0 1-1 1H2a1 1 0 0 1-1-1V4Zm3 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0Zm4.5 2L7 3H6v3h2.5ZM10 6h1.5L10 3H8.5L10 6Zm3.25 0H13l-1.5-3h1.75a.75.75 0 0 1 .75.75v1.5a.75.75 0 0 1-.75.75Z"/>`),
		g.Group(children),
	)
}