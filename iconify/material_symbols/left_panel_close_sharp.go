package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftPanelCloseSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.5 16V8l-4 4l4 4ZM10 19h9V5h-9v14Zm-7 2V3h18v18H3Z"/>`),
		g.Group(children),
	)
}