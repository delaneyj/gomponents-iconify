package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalHospitalRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 13.5v2q0 .625.438 1.063T12 17q.625 0 1.063-.438T13.5 15.5v-2h2q.625 0 1.063-.438T17 12q0-.625-.438-1.063T15.5 10.5h-2v-2q0-.625-.438-1.063T12 7q-.625 0-1.063.438T10.5 8.5v2h-2q-.625 0-1.063.438T7 12q0 .625.438 1.063T8.5 13.5h2ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}