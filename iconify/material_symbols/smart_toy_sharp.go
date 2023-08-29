package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartToySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15q-1.25 0-2.125-.875T1 12q0-1.25.875-2.125T4 9V5h5q0-1.25.875-2.125T12 2q1.25 0 2.125.875T15 5h5v4q1.25 0 2.125.875T23 12q0 1.25-.875 2.125T20 15v6H4v-6Zm5-2q.625 0 1.063-.438T10.5 11.5q0-.625-.438-1.063T9 10q-.625 0-1.063.438T7.5 11.5q0 .625.438 1.063T9 13Zm6 0q.625 0 1.063-.438T16.5 11.5q0-.625-.438-1.063T15 10q-.625 0-1.063.438T13.5 11.5q0 .625.438 1.063T15 13Zm-7 4h8v-2H8v2Z"/>`),
		g.Group(children),
	)
}