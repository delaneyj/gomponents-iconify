package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyChangingStationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V12l1.575-4.625q.275-.85 1.1-1.2t1.625 0L11.45 8H14v2h-3L8.3 8.825L7 12.75V22H3ZM8 5q-.825 0-1.413-.588T6 3q0-.825.588-1.413T8 1q.825 0 1.413.588T10 3q0 .825-.588 1.413T8 5Zm1 14v-2h12v2H9Zm10.5-3q-.625 0-1.063-.438T18 14.5q0-.625.438-1.063T19.5 13q.625 0 1.063.438T21 14.5q0 .625-.438 1.063T19.5 16ZM11 16v-3H9v-2h4v2h2v-2h2v5h-6Z"/>`),
		g.Group(children),
	)
}