package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeMaxDots(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20v-1H5q-1.65 0-2.825-1.175T1 15V9q0-1.65 1.175-2.825T5 5h14q1.65 0 2.825 1.175T23 9v6q0 1.65-1.175 2.825T19 19h-2v1H7Zm3.5-7q.425 0 .713-.288T11.5 12q0-.425-.288-.713T10.5 11q-.425 0-.713.288T9.5 12q0 .425.288.713T10.5 13Zm-3 0q.425 0 .713-.288T8.5 12q0-.425-.288-.713T7.5 11q-.425 0-.713.288T6.5 12q0 .425.288.713T7.5 13Zm6 0q.425 0 .713-.288T14.5 12q0-.425-.288-.713T13.5 11q-.425 0-.713.288T12.5 12q0 .425.288.713T13.5 13Zm3 0q.425 0 .713-.288T17.5 12q0-.425-.288-.713T16.5 11q-.425 0-.713.288T15.5 12q0 .425.288.713T16.5 13Z"/>`),
		g.Group(children),
	)
}