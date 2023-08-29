package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongDownE(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.998 1.02h-6v6h2.001l.013 12.145l-1.844-1.834l-1.41 1.418l4.254 4.23l4.23-4.254l-1.417-1.41l-1.813 1.823l-.013-12.118h1.999v-6Zm-4 2h2v2h-2v-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}