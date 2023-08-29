package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardsVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 2h14a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1m1 2v8h12V4H6m14 13a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16v1m0 4a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-1h16v1Z"/>`),
		g.Group(children),
	)
}