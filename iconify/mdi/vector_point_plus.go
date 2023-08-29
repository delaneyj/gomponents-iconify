package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorPointPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 9v6h6V9H9m2 2h2v2h-2v-2m7 4v3h-3v2h3v3h2v-3h3v-2h-3v-3h-2Z"/>`),
		g.Group(children),
	)
}