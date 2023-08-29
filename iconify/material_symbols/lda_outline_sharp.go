package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LdaOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-3.425L5 16.05V11h6V8H8V2h8v6h-3v3h6v5.05l-6 2.525V22h-2ZM10 6h4V4h-4v2Zm1 10.4V13H7v1.725l4 1.675Zm2 0l4-1.675V13h-4v3.4ZM10 6V4v2Z"/>`),
		g.Group(children),
	)
}