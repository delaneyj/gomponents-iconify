package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AodTabletSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20V4h22v16H1Zm5-2h12V6H6v12Zm2-6.5V10h8v1.5H8Zm1 3V13h6v1.5H9Z"/>`),
		g.Group(children),
	)
}