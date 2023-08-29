package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 11q-1.45 0-2.475-1.025T4 7.5q0-1.45 1.025-2.475T7.5 4q1.45 0 2.475 1.025T11 7.5q0 1.45-1.025 2.475T7.5 11Zm0-2q.625 0 1.063-.438T9 7.5q0-.625-.438-1.063T7.5 6q-.625 0-1.063.438T6 7.5q0 .625.438 1.063T7.5 9Zm9 11q-1.45 0-2.475-1.025T13 16.5q0-1.45 1.025-2.475T16.5 13q1.45 0 2.475 1.025T20 16.5q0 1.45-1.025 2.475T16.5 20Zm0-2q.625 0 1.063-.438T18 16.5q0-.625-.438-1.063T16.5 15q-.625 0-1.063.438T15 16.5q0 .625.438 1.063T16.5 18ZM4.7 19.3q-.275-.275-.275-.7t.275-.7L17.9 4.7q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7L6.1 19.3q-.275.275-.7.275t-.7-.275Z"/>`),
		g.Group(children),
	)
}