package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 19.15L14.85 12H20V8h-9.15l-4-4H22v15.15ZM4 12h5.15l-4-4H4v4Zm16.45 11.3l-3.3-3.3H2V4h2l2 2H3.15L.65 3.5l1.425-1.425l19.8 19.8L20.45 23.3Z"/>`),
		g.Group(children),
	)
}