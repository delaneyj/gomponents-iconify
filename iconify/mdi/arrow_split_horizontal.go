package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSplitHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 18h3v-3H2v-2h20v2h-9v3h3l-4 4l-4-4m4-16L8 6h3v3H2v2h20V9h-9V6h3l-4-4Z"/>`),
		g.Group(children),
	)
}