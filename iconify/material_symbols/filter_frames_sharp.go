package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterFramesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V4h6l4-4l4 4h6v18H2Zm2-2h16V6H4v14Zm2-2V8h12v10H6Z"/>`),
		g.Group(children),
	)
}