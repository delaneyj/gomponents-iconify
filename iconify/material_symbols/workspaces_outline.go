package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WorkspacesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-1.65 0-2.825-1.175T2 17q0-1.65 1.175-2.825T6 13q1.65 0 2.825 1.175T10 17q0 1.65-1.175 2.825T6 21Zm12 0q-1.65 0-2.825-1.175T14 17q0-1.65 1.175-2.825T18 13q1.65 0 2.825 1.175T22 17q0 1.65-1.175 2.825T18 21ZM6 19q.825 0 1.413-.588T8 17q0-.825-.588-1.413T6 15q-.825 0-1.413.588T4 17q0 .825.588 1.413T6 19Zm12 0q.825 0 1.413-.588T20 17q0-.825-.588-1.413T18 15q-.825 0-1.413.588T16 17q0 .825.588 1.413T18 19Zm-6-8q-1.65 0-2.825-1.175T8 7q0-1.65 1.175-2.825T12 3q1.65 0 2.825 1.175T16 7q0 1.65-1.175 2.825T12 11Zm0-2q.825 0 1.413-.588T14 7q0-.825-.588-1.413T12 5q-.825 0-1.413.588T10 7q0 .825.588 1.413T12 9Zm0-2Zm6 10ZM6 17Z"/>`),
		g.Group(children),
	)
}