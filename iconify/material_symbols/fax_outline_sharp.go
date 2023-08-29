package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaxOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V8h6v1v-5h10v5h4v11H8v1H2Zm2-2h2v-9H4v9Zm6-10h6V6h-6v3Zm-2 9h12v-7H8v7Zm7-4q.425 0 .713-.288T16 13q0-.425-.288-.713T15 12q-.425 0-.713.288T14 13q0 .425.288.713T15 14Zm3 0q.425 0 .713-.288T19 13q0-.425-.288-.713T18 12q-.425 0-.713.288T17 13q0 .425.288.713T18 14Zm-3 3q.425 0 .713-.288T16 16q0-.425-.288-.713T15 15q-.425 0-.713.288T14 16q0 .425.288.713T15 17Zm3 0q.425 0 .713-.288T19 16q0-.425-.288-.713T18 15q-.425 0-.713.288T17 16q0 .425.288.713T18 17Zm-9 0h4v-5H9v5Zm-1 1v-7v7Z"/>`),
		g.Group(children),
	)
}