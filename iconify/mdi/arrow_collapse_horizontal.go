package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCollapseHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 20V4h2.03v16H13m-3 0V4h2.03v16H10M5 8l4.03 4L5 16v-3H2v-2h3V8m15 8l-4-4l4-4v3h3v2h-3v3Z"/>`),
		g.Group(children),
	)
}