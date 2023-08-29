package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalLibraryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22.5q-1.8-1.7-4.125-2.6T3 19V8q2.525 0 4.85.913T12 11.55q1.825-1.725 4.15-2.638T21 8v11q-2.575 0-4.888.9T12 22.5Zm0-2.6q1.575-1.175 3.35-1.875T19 17.1v-6.9q-1.825.325-3.588 1.313T12 14.15q-1.65-1.65-3.413-2.638T5 10.2v6.9q1.875.225 3.65.925T12 19.9ZM12 9q-1.65 0-2.825-1.175T8 5q0-1.65 1.175-2.825T12 1q1.65 0 2.825 1.175T16 5q0 1.65-1.175 2.825T12 9Zm0-2q.825 0 1.413-.588T14 5q0-.825-.588-1.413T12 3q-.825 0-1.413.588T10 5q0 .825.588 1.413T12 7Zm0-2Zm0 9.15Z"/>`),
		g.Group(children),
	)
}