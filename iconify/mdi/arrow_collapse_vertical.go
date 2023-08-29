package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCollapseVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 12h16v2H4v-2m0-3h16v2H4V9m12-5l-4 4l-4-4h3V1h2v3h3M8 19l4-4l4 4h-3v3h-2v-3H8Z"/>`),
		g.Group(children),
	)
}