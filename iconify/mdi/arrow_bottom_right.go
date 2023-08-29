package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 6.41L6.41 5L17 15.59V9h2v10H9v-2h6.59L5 6.41Z"/>`),
		g.Group(children),
	)
}