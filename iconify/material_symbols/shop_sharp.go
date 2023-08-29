package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V6h6V2h8v4h6v15H2Zm8-15h4V4h-4v2Zm-.5 12l7-4.5l-7-4.5v9Z"/>`),
		g.Group(children),
	)
}