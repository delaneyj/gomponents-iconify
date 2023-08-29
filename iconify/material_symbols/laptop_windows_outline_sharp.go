package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopWindowsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 20v-2h4v-1H2V3h20v14h-2v1h4v2H0Zm4-5h16V5H4v10Zm0 0V5v10Z"/>`),
		g.Group(children),
	)
}