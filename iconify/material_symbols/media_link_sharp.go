package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaLinkSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.5 19.5l4-2.5l-4-2.5v5ZM13 7.75h4v-1.5h-4v1.5ZM1 23V11h16v12H1Zm17.5-9V9.6q.675-.4 1.088-1.087T20 7q0-1.25-.875-2.125T17 4h-1.25v1.5H17q.625 0 1.063.438T18.5 7q0 .625-.438 1.063T17 8.5h-1.25v1h-1.5v-1H13q-.625 0-1.063-.438T11.5 7q0-.625.438-1.063T13 5.5h1.25V4H13q-1.25 0-2.125.875T10 7q0 .8.375 1.438T11.35 9.5H7V1h16v13h-4.5Z"/>`),
		g.Group(children),
	)
}