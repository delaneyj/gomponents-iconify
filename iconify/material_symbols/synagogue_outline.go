package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SynagogueOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V7q0-1.25.875-2.125T4 4q1.25 0 2.125.875T7 7v.3L12 3l5 4.3V7q0-1.25.875-2.125T20 4q1.25 0 2.125.875T23 7v14H13v-5q0-.425-.288-.713T12 15q-.425 0-.713.288T11 16v5H1ZM19 8h2V7q0-.425-.288-.713T20 6q-.425 0-.713.288T19 7v1ZM3 8h2V7q0-.425-.288-.713T4 6q-.425 0-.713.288T3 7v1Zm0 11h2v-9H3v9Zm4 0h2v-3q0-1.25.875-2.125T12 13q1.25 0 2.125.875T15 16v3h2V9.925l-5-4.3l-5 4.3V19Zm12 0h2v-9h-2v9Zm-7-7.5q-.625 0-1.063-.438T10.5 10q0-.625.438-1.063T12 8.5q.625 0 1.063.438T13.5 10q0 .625-.438 1.063T12 11.5Z"/>`),
		g.Group(children),
	)
}