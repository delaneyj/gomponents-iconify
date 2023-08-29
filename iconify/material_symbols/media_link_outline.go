package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaLinkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 19.5v-5l4 2.5l-4 2.5ZM13 10q-1.25 0-2.125-.875T10 7q0-1.25.875-2.125T13 4h1.25v1.5H13q-.625 0-1.063.438T11.5 7q0 .625.438 1.063T13 8.5h1.25V10H13Zm2.75 0V8.5H17q.625 0 1.063-.438T18.5 7q0-.625-.438-1.063T17 5.5h-1.25V4H17q1.25 0 2.125.875T20 7q0 1.25-.875 2.125T17 10h-1.25ZM13 7.75v-1.5h4v1.5h-4ZM16.1 14v-2H21V3H9v9H7V3q0-.825.588-1.413T9 1h12q.825 0 1.413.588T23 3v9q0 .825-.588 1.413T21 14h-4.9ZM3 23q-.825 0-1.413-.588T1 21v-8q0-.825.588-1.413T3 11h12q.825 0 1.413.588T17 13v8q0 .825-.588 1.413T15 23H3Zm0-2h12v-8H3v8ZM15 7.5ZM9 17Z"/>`),
		g.Group(children),
	)
}