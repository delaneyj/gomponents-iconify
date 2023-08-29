package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19q-.425 0-.713-.288T1 18V5q0-.425.288-.713T2 4q.425 0 .713.288T3 5v9h8V8q0-.825.588-1.413T13 6h6q1.65 0 2.825 1.175T23 10v8q0 .425-.288.713T22 19q-.425 0-.713-.288T21 18v-2H3v2q0 .425-.288.713T2 19Zm5-6q-1.25 0-2.125-.875T4 10q0-1.25.875-2.125T7 7q1.25 0 2.125.875T10 10q0 1.25-.875 2.125T7 13Zm6 1h8v-4q0-.825-.588-1.413T19 8h-6v6Zm-6-3q.425 0 .713-.288T8 10q0-.425-.288-.713T7 9q-.425 0-.713.288T6 10q0 .425.288.713T7 11Zm0-1Zm6-2v6v-6Z"/>`),
		g.Group(children),
	)
}