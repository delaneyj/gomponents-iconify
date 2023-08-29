package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomPanelOpenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 11.5h8l-4-4l-4 4ZM5 14h14V5H5v9Zm-2 7V3h18v18H3Z"/>`),
		g.Group(children),
	)
}