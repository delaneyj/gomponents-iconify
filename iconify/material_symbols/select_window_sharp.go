package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectWindowSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V9h4V2h16v13h-4v7H2Zm2-2h12v-7H4v7Zm14-7h2V6H8v3h10v4Z"/>`),
		g.Group(children),
	)
}