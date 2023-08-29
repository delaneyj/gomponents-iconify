package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireHydrantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-2h2v-3H5q-.825 0-1.413-.588T3 15v-2q0-.825.588-1.413T5 11h1V8H4V6h2.35q.625-1.75 2.163-2.875T12 2q1.95 0 3.488 1.125T17.65 6H20v2h-2v3h1q.825 0 1.413.588T21 13v2q0 .825-.588 1.413T19 17h-1v3h2v2H4ZM8.525 6h6.95q-.525-.9-1.437-1.45T12 4q-1.125 0-2.038.55T8.526 6ZM8 20h8v-5h3v-2h-3V8H8v5H5v2h3v5Zm4-2.5q1.45 0 2.475-1.025T15.5 14q0-1.45-1.025-2.475T12 10.5q-1.45 0-2.475 1.025T8.5 14q0 1.45 1.025 2.475T12 17.5Zm0-2q-.625 0-1.063-.438T10.5 14q0-.625.438-1.063T12 12.5q.625 0 1.063.438T13.5 14q0 .625-.438 1.063T12 15.5Zm0-1.5Zm0-6Z"/>`),
		g.Group(children),
	)
}