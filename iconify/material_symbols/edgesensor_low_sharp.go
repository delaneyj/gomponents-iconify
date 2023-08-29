package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgesensorLowSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 14V7h2v7H2Zm18 3v-7h2v7h-2ZM6 22V2h12v20H6Zm2-5h8V7H8v10Z"/>`),
		g.Group(children),
	)
}