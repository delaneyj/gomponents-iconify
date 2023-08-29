package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FrontHandOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 1q.425 0 .713.288T14 2v9q0 .425-.288.713T13 12q-.425 0-.713-.288T12 11V2q0-.425.288-.713T13 1ZM9 2q.425 0 .713.288T10 3v8q0 .425-.288.713T9 12q-.425 0-.713-.288T8 11V3q0-.425.288-.713T9 2Zm3.5 21q-3.55 0-6.025-2.475T4 14.5V5q0-.425.288-.713T5 4q.425 0 .713.288T6 5v9.5q0 2.725 1.888 4.612T12.5 21q2.725 0 4.612-1.888T19 14.5V11q-.425 0-.713.288T18 12v3q0 .425-.288.713T17 16h-2q-.825 0-1.413.588T13 18v1q0 .425-.288.713T12 20q-.425 0-.713-.288T11 19v-1q0-1.65 1.175-2.825T15 14h1V4q0-.425.288-.713T17 3q.425 0 .713.288T18 4v5.175q.25-.075.488-.125T19 9h1q.425 0 .713.288T21 10v4.5q0 3.55-2.475 6.025T12.5 23Zm1-8Z"/>`),
		g.Group(children),
	)
}