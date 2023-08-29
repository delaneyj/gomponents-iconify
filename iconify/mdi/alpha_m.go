package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaM(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7a2 2 0 0 0-2 2v8h2V9h2v7h2V9h2v8h2V9a2 2 0 0 0-2-2H9Z"/>`),
		g.Group(children),
	)
}