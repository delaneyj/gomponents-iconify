package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ethernet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 .5v20h7v3h2v-3h7V.5H4Zm14 2H6v6h4v6h4v-6h4v-6Zm-12 16v-8h2v6h8v-6h2v8H6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}