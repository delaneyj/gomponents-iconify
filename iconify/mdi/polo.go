package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 9.5a2.5 2.5 0 0 1-5 0a2.5 2.5 0 0 1 5 0M11 17V3H8v14H2l4 4h7v-4h-2m5 0h-2v4h2v-4Z"/>`),
		g.Group(children),
	)
}