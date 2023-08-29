package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopWindowsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-2h2v-2H2V3h20v14h-8v2h2v2H8Zm-4-6h16V5H4v10Zm0 0V5v10Z"/>`),
		g.Group(children),
	)
}