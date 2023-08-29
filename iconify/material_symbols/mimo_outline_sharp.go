package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MimoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2l1-1H2V3h20v15h-5l1 1v2H6Zm-2-5h16V5H4v11Zm0 0V5v11Z"/>`),
		g.Group(children),
	)
}