package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphaD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7v10h4a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2H9m2 2h2v6h-2V9Z"/>`),
		g.Group(children),
	)
}