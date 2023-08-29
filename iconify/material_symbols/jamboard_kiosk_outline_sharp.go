package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JamboardKioskOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2h5v-3H2V3h20v13h-9v3h5v2H6Zm-2-7h16V5H4v9Zm0 0V5v9Z"/>`),
		g.Group(children),
	)
}