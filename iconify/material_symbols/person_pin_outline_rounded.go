package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonPinOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20H5q-.825 0-1.413-.588T3 18V4q0-.825.588-1.413T5 2h14q.825 0 1.413.588T21 4v14q0 .825-.588 1.413T19 20h-4l-2.3 2.3q-.15.15-.325.213t-.375.062q-.2 0-.375-.063T11.3 22.3L9 20Zm-4-3.15q1.35-1.325 3.138-2.087T12 14q2.075 0 3.863.763T19 16.85V4H5v12.85ZM12 12q1.45 0 2.475-1.025T15.5 8.5q0-1.45-1.025-2.475T12 5q-1.45 0-2.475 1.025T8.5 8.5q0 1.45 1.025 2.475T12 12Zm-5 6h10v-.25q-1.05-.875-2.325-1.313T12 16q-1.4 0-2.675.438T7 17.75V18Zm5-8q-.625 0-1.063-.438T10.5 8.5q0-.625.438-1.063T12 7q.625 0 1.063.438T13.5 8.5q0 .625-.438 1.063T12 10Zm0 .425Z"/>`),
		g.Group(children),
	)
}