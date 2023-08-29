package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopMacOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 21v-1l2-2H2V3h20v15h-8l2 2v1H8Zm-4-8h16V5H4v8Zm0 0V5v8Z"/>`),
		g.Group(children),
	)
}