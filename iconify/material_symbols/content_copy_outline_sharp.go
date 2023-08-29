package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContentCopyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18V2h13v16H7Zm2-2h9V4H9v12Zm-6 6V6h2v14h11v2H3Zm6-6V4v12Z"/>`),
		g.Group(children),
	)
}