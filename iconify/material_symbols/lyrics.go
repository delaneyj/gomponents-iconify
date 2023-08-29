package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lyrics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h4v-2H6v2Zm13-2q-1.25 0-2.125-.875T16 9q0-1.25.875-2.125T19 6q.275 0 .513.05t.487.125V1h4v2h-2v6q0 1.25-.875 2.125T19 12ZM6 11h7V9H6v2Zm0-3h7V6H6v2Zm0 10l-4 4V4q0-.825.588-1.413T4 2h11q.825 0 1.413.588T17 4v.425q-1.375.6-2.188 1.838T14 9q0 1.5.813 2.738T17 13.575V16q0 .825-.588 1.413T15 18H6Z"/>`),
		g.Group(children),
	)
}