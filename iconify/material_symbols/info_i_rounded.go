package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfoIRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 7q-.825 0-1.413-.588T10 5q0-.825.588-1.413T12 3q.825 0 1.413.588T14 5q0 .825-.588 1.413T12 7Zm0 14q-.625 0-1.063-.438T10.5 19.5v-9q0-.625.438-1.063T12 9q.625 0 1.063.438T13.5 10.5v9q0 .625-.438 1.063T12 21Z"/>`),
		g.Group(children),
	)
}