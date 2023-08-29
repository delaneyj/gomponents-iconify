package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DomainVerificationOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h16V8H4v10Zm6.95-1.45L7.4 13l1.45-1.45l2.1 2.1l4.2-4.2l1.45 1.45l-5.65 5.65ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}