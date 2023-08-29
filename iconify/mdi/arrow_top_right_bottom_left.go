package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopRightBottomLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21H3v-8h2v4.59L17.59 5H13V3h8v8h-2V6.41L6.41 19H11v2Z"/>`),
		g.Group(children),
	)
}