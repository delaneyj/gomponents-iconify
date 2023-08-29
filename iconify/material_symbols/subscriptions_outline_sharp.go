package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubscriptionsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V8h20v14H2Zm2-2h16V10H4v10Zm6-1l6-4l-6-4v8ZM4 7V5h16v2H4Zm3-3V2h10v2H7ZM4 20V10v10Z"/>`),
		g.Group(children),
	)
}