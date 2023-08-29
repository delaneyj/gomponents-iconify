package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccountBoxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17.85q1.35-1.325 3.138-2.087T12 15q2.075 0 3.863.763T19 17.85V5H5v12.85ZM12 13q1.45 0 2.475-1.025T15.5 9.5q0-1.45-1.025-2.475T12 6q-1.45 0-2.475 1.025T8.5 9.5q0 1.45 1.025 2.475T12 13Zm-7 8q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm2-2h10v-.25q-1.05-.875-2.325-1.313T12 17q-1.4 0-2.675.438T7 18.75V19Zm5-8q-.625 0-1.063-.438T10.5 9.5q0-.625.438-1.063T12 8q.625 0 1.063.438T13.5 9.5q0 .625-.438 1.063T12 11Zm0 .425Z"/>`),
		g.Group(children),
	)
}