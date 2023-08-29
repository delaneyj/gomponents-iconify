package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireHydrantRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.425 0-.713-.288T4 21q0-.425.288-.713T5 20h1v-3H5q-.825 0-1.413-.588T3 15v-2q0-.825.588-1.413T5 11h1V8H5q-.425 0-.713-.288T4 7q0-.425.288-.713T5 6h1.35q.625-1.75 2.163-2.875T12 2q1.95 0 3.488 1.125T17.65 6H19q.425 0 .713.288T20 7q0 .425-.288.713T19 8h-1v3h1q.825 0 1.413.588T21 13v2q0 .825-.588 1.413T19 17h-1v3h1q.425 0 .713.288T20 21q0 .425-.288.713T19 22H5Zm7-4.5q1.45 0 2.475-1.025T15.5 14q0-1.45-1.025-2.475T12 10.5q-1.45 0-2.475 1.025T8.5 14q0 1.45 1.025 2.475T12 17.5Zm0-2q-.625 0-1.063-.438T10.5 14q0-.625.438-1.063T12 12.5q.625 0 1.063.438T13.5 14q0 .625-.438 1.063T12 15.5Z"/>`),
		g.Group(children),
	)
}