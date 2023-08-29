package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocamOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h16v6.5l4-4v11l-4-4V20H2Zm2-2h12V6H4v12ZM4 6v12V6Z"/>`),
		g.Group(children),
	)
}