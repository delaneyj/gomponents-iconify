package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AodTabletOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20V4h22v16H1ZM4 6H3v12h1V6Zm2 12h12V6H6v12ZM20 6v12h1V6h-1Zm0 0h1h-1ZM4 6H3h1Zm4 5.5V10h8v1.5H8Zm1 3V13h6v1.5H9Z"/>`),
		g.Group(children),
	)
}