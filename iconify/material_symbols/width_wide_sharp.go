package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidthWideSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h2V6H4v12Zm14 0h2V6h-2v12Z"/>`),
		g.Group(children),
	)
}