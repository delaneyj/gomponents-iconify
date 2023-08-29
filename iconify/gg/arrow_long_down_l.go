package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongDownL(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.998.972v2h3l-1 .001l.014 16.24l-1.844-1.834l-1.41 1.418l4.254 4.23l4.23-4.254l-1.417-1.41l-1.813 1.823l-.014-16.214h2v-2h-6Z"/>`),
		g.Group(children),
	)
}