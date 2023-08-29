package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieInfoOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V3h20v18H2Zm2-2h2v-2H4v2Zm14 0h2v-2h-2v2Zm-7-2h2v-6h-2v6Zm-7-2h2v-2H4v2Zm14 0h2v-2h-2v2ZM4 11h2V9H4v2Zm14 0h2V9h-2v2Zm-6-2q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7q-.425 0-.713.288T11 8q0 .425.288.713T12 9ZM4 7h2V5H4v2Zm14 0h2V5h-2v2ZM8 19h8V5H8v14ZM8 5h8h-8Z"/>`),
		g.Group(children),
	)
}