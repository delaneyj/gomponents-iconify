package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTravelOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19h16v-2H4v2Zm-2 2V6h5V2h10v4h5v15H2Zm2-7h16V8h-3v2h-2V8H9v2H7V8H4v6Zm5-8h6V4H9v2ZM4 19V8v2v-2v2v-2v11Z"/>`),
		g.Group(children),
	)
}