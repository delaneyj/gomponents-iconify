package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeftBottomRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 21h8v-8h-2v4.59L6.41 5H11V3H3v8h2V6.41L17.59 19H13v2Z"/>`),
		g.Group(children),
	)
}