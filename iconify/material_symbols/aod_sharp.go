package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AodSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14.5V13h6v1.5H9Zm-1-3V10h8v1.5H8ZM5 23V1h14v22H5Zm2-5h10V6H7v12Z"/>`),
		g.Group(children),
	)
}