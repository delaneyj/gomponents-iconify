package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToolsLadderSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.7 21l5-18h2l-.85 3h5.625l.825-3h2l-5 18h-2l.85-3H7.525L6.7 21h-2Zm4.775-10h5.6l.825-3h-5.6l-.825 3ZM8.1 16h5.6l.825-3h-5.6L8.1 16Z"/>`),
		g.Group(children),
	)
}