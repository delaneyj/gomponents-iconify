package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgesensorHighSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 17v-7h2v7H0Zm3-3V7h2v7H3Zm3 8V2h12v20H6Zm13-5v-7h2v7h-2Zm3-3V7h2v7h-2ZM8 17h8V7H8v10Z"/>`),
		g.Group(children),
	)
}