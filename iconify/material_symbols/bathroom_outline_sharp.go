package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BathroomOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 18q.425 0 .713-.288T10 17q0-.425-.288-.713T9 16q-.425 0-.713.288T8 17q0 .425.288.713T9 18Zm3 0q.425 0 .713-.288T13 17q0-.425-.288-.713T12 16q-.425 0-.713.288T11 17q0 .425.288.713T12 18Zm3 0q.425 0 .713-.288T16 17q0-.425-.288-.713T15 16q-.425 0-.713.288T14 17q0 .425.288.713T15 18Zm-6-3q.425 0 .713-.288T10 14q0-.425-.288-.713T9 13q-.425 0-.713.288T8 14q0 .425.288.713T9 15Zm3 0q.425 0 .713-.288T13 14q0-.425-.288-.713T12 13q-.425 0-.713.288T11 14q0 .425.288.713T12 15Zm3 0q.425 0 .713-.288T16 14q0-.425-.288-.713T15 13q-.425 0-.713.288T14 14q0 .425.288.713T15 15Zm-8-3h10v-1q0-2.075-1.463-3.538T12 6Q9.925 6 8.462 7.463T7 11v1Zm1.55-1.5q.2-1.275 1.163-2.138T12 7.5q1.325 0 2.288.863T15.45 10.5h-6.9ZM2 22V2h20v20H2Zm2-2h16V4H4v16Zm0 0V4v16Z"/>`),
		g.Group(children),
	)
}