package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThirtyFpsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19H3.5q-.625 0-1.063-.438T2 17.5q0-.625.438-1.063T3.5 16H8v-2.5H3.5q-.625 0-1.063-.438T2 12q0-.625.438-1.063T3.5 10.5H8V8H3.5q-.625 0-1.063-.438T2 6.5q0-.625.438-1.063T3.5 5H8q1.25 0 2.125.875T11 8v1.9q0 .875-.613 1.488T8.9 12q.875 0 1.488.613T11 14.1V16q0 1.25-.875 2.125T8 19Zm7-3h4V8h-4v8Zm0 3q-1.25 0-2.125-.875T12 16V8q0-1.25.875-2.125T15 5h4q1.25 0 2.125.875T22 8v8q0 1.25-.875 2.125T19 19h-4Z"/>`),
		g.Group(children),
	)
}