package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V2h20v20H2Zm2-2h16V4H4v16Zm5-6q-.425 0-.713-.288T8 13q0-.425.288-.713T9 12q.425 0 .713.288T10 13q0 .425-.288.713T9 14Zm6 0q-.425 0-.713-.288T14 13q0-.425.288-.713T15 12q.425 0 .713.288T16 13q0 .425-.288.713T15 14ZM5 11.1v7.4h2v-2h10v2h2v-7.4l-1.925-5.6H6.925L5 11.1Zm2.65-1.6l.7-2h7.3l.7 2h-8.7ZM4 4v16V4Zm3 10.5v-3h10v3H7Z"/>`),
		g.Group(children),
	)
}