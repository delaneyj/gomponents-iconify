package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalTaxiOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 19v1q0 .425-.288.713T5 21H4q-.425 0-.713-.288T3 20v-8l2.1-6q.15-.45.537-.725T6.5 5H9V3h6v2h2.5q.475 0 .863.275T18.9 6l2.1 6v8q0 .425-.287.713T20 21h-1q-.425 0-.713-.288T18 20v-1H6Zm-.2-9h12.4l-1.05-3H6.85L5.8 10ZM5 12v5v-5Zm2.5 4q.625 0 1.063-.438T9 14.5q0-.625-.438-1.063T7.5 13q-.625 0-1.063.438T6 14.5q0 .625.438 1.063T7.5 16Zm9 0q.625 0 1.063-.438T18 14.5q0-.625-.438-1.063T16.5 13q-.625 0-1.063.438T15 14.5q0 .625.438 1.063T16.5 16ZM5 17h14v-5H5v5Z"/>`),
		g.Group(children),
	)
}