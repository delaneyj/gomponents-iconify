package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DraftsOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l10 5.975V21H2V6.975L12 1Zm0 11.65L19.8 8L12 3.35L4.2 8l7.8 4.65ZM12 15l-8-4.8V19h16v-8.8L12 15Zm0 4h8H4h8Z"/>`),
		g.Group(children),
	)
}