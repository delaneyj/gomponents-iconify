package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LivingOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V2h20v20H2Zm2-2h16V4H4v16Zm1-2h14v-7.65h-1.25V6H6.25v4.35H5V18Zm1.5-1.5v-5h2v3h7v-3h2v5h-11ZM10 13v-2.95H7.75V7.5h8.5v2.55H14V13h-4Zm-6 7V4v16Z"/>`),
		g.Group(children),
	)
}