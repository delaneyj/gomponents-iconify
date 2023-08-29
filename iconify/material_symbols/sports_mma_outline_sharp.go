package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMmaOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21v-4h10v4H7Zm-1-5l-1-5.05V3h12v4h2v3.95L18 16H6Zm1.65-2h8.7l.65-3.4V10h-2V5H7v5.6l.65 3.4ZM8 10h6V7H8v3Zm4-.5Z"/>`),
		g.Group(children),
	)
}