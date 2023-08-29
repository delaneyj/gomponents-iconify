package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V7H5v12Zm7-2q-2.05 0-3.663-1.113T6 13q.725-1.775 2.337-2.888T12 9q2.05 0 3.663 1.113T18 13q-.725 1.775-2.337 2.888T12 17Zm0-2.5q-.625 0-1.063-.438T10.5 13q0-.625.438-1.063T12 11.5q.625 0 1.063.438T13.5 13q0 .625-.438 1.063T12 14.5Zm0 1q1.05 0 1.775-.725T14.5 13q0-1.05-.725-1.775T12 10.5q-1.05 0-1.775.725T9.5 13q0 1.05.725 1.775T12 15.5Z"/>`),
		g.Group(children),
	)
}