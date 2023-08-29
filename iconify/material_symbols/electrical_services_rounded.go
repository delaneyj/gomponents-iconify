package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricalServicesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 15v-2h2q.425 0 .713.288T21 14q0 .425-.288.713T20 15h-2Zm0 4v-2h2q.425 0 .713.288T21 18q0 .425-.288.713T20 19h-2Zm-4 1q-.825 0-1.413-.588T12 18h-2v-4h2q0-.825.588-1.413T14 12h2q.425 0 .713.288T17 13v6q0 .425-.288.713T16 20h-2Zm-7-3q-1.65 0-2.825-1.175T3 13q0-1.65 1.175-2.825T7 9h1.5q.625 0 1.063-.438T10 7.5q0-.625-.438-1.063T8.5 6H5q-.425 0-.713-.288T4 5q0-.425.288-.713T5 4h3.5q1.45 0 2.475 1.025T12 7.5q0 1.45-1.025 2.475T8.5 11H7q-.825 0-1.413.588T5 13q0 .825.588 1.413T7 15h2v2H7Z"/>`),
		g.Group(children),
	)
}