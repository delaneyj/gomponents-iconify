package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 13q.425 0 .713-.288T10.5 12q0-.425-.288-.713T9.5 11q-.425 0-.713.288T8.5 12q0 .425.288.713T9.5 13Zm3.5 0q.425 0 .713-.288T14 12q0-.425-.288-.713T13 11q-.425 0-.713.288T12 12q0 .425.288.713T13 13Zm3.5 0q.425 0 .713-.288T17.5 12q0-.425-.288-.713T16.5 11q-.425 0-.713.288T15.5 12q0 .425.288.713T16.5 13Zm-8.55 6L3 12l4.95-7H21v14H7.95Z"/>`),
		g.Group(children),
	)
}