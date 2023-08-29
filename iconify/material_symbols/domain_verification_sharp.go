package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainVerificationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.95 16.55L7.4 13l1.45-1.45l2.1 2.1l4.2-4.2l1.45 1.45l-5.65 5.65ZM4 8V6h16v2H4ZM2 20h20V4H2v16Z"/>`),
		g.Group(children),
	)
}