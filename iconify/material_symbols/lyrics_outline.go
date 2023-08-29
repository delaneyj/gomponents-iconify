package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LyricsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 16V4v12Zm-2 6V4q0-.825.588-1.413T4 2h11q.825 0 1.413.588T17 4v.425q-.6.275-1.1.675T15 6V4H4v12h11v-4q.4.5.9.9t1.1.675V16q0 .825-.588 1.413T15 18H6l-4 4Zm4-8h4v-2H6v2Zm13-2q-1.25 0-2.125-.875T16 9q0-1.25.875-2.125T19 6q.275 0 .525.05t.475.125V1h4v2h-2v6q0 1.25-.875 2.125T19 12ZM6 11h7V9H6v2Zm0-3h7V6H6v2Z"/>`),
		g.Group(children),
	)
}