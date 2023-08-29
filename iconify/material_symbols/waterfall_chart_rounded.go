package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterfallChartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.5 20q-.625 0-1.063-.438T18 18.5v-13q0-.625.438-1.063T19.5 4q.625 0 1.063.438T21 5.5v13q0 .625-.438 1.063T19.5 20Zm-15 0q-.625 0-1.063-.438T3 18.5v-4q0-.625.438-1.063T4.5 13q.625 0 1.063.438T6 14.5v4q0 .625-.438 1.063T4.5 20Zm11-13q-.625 0-1.063-.438T14 5.5q0-.625.438-1.063T15.5 4q.625 0 1.063.438T17 5.5q0 .625-.438 1.063T15.5 7Zm-4 2q-.625 0-1.063-.438T10 7.5v-1q0-.625.438-1.063T11.5 5q.625 0 1.063.438T13 6.5v1q0 .625-.438 1.063T11.5 9Zm-3 5q-.625 0-1.063-.438T7 12.5v-1q0-.625.438-1.063T8.5 10q.625 0 1.063.438T10 11.5v1q0 .625-.438 1.063T8.5 14Z"/>`),
		g.Group(children),
	)
}