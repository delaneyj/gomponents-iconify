package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpokeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 11q-1.65 0-2.825-1.175T8 7q0-1.65 1.175-2.825T12 3q1.65 0 2.825 1.175T16 7q0 1.65-1.175 2.825T12 11Zm0-2q.825 0 1.413-.588T14 7q0-.825-.588-1.413T12 5q-.825 0-1.413.588T10 7q0 .825.588 1.413T12 9ZM7 21q-1.65 0-2.825-1.175T3 17q0-1.65 1.175-2.825T7 13q1.65 0 2.825 1.175T11 17q0 1.65-1.175 2.825T7 21Zm0-2q.825 0 1.413-.588T9 17q0-.825-.588-1.413T7 15q-.825 0-1.413.588T5 17q0 .825.588 1.413T7 19Zm10 2q-1.65 0-2.825-1.175T13 17q0-1.65 1.175-2.825T17 13q1.65 0 2.825 1.175T21 17q0 1.65-1.175 2.825T17 21Zm0-2q.825 0 1.413-.588T19 17q0-.825-.588-1.413T17 15q-.825 0-1.413.588T15 17q0 .825.588 1.413T17 19ZM12 7ZM7 17Zm10 0Z"/>`),
		g.Group(children),
	)
}