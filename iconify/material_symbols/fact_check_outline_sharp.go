package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactCheckOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17h5v-2H5v2Zm9.55-2l4.95-4.95l-1.425-1.425l-3.525 3.55l-1.425-1.425l-1.4 1.425L14.55 15ZM5 13h5v-2H5v2Zm0-4h5V7H5v2ZM2 21V3h20v18H2Zm2-2h16V5H4v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}