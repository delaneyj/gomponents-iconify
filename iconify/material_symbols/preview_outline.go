package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PreviewOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V7H5v12Zm7-2q-2.05 0-3.663-1.113T6 13q.725-1.775 2.337-2.888T12 9q2.05 0 3.663 1.113T18 13q-.725 1.775-2.337 2.888T12 17Zm0-1.5q1.4 0 2.55-.663T16.35 13q-.65-1.175-1.8-1.838T12 10.5q-1.4 0-2.55.663T7.65 13q.65 1.175 1.8 1.838T12 15.5Zm0-2.5Zm0 1.5q.625 0 1.063-.438T13.5 13q0-.625-.438-1.063T12 11.5q-.625 0-1.063.438T10.5 13q0 .625.438 1.063T12 14.5Z"/>`),
		g.Group(children),
	)
}