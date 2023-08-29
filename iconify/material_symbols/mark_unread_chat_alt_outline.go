package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkUnreadChatAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h8v-2H6v2Zm0-3h12V9H6v2Zm0-3h12v-.1q-.925-.2-1.688-.688T15 6H6v2ZM2 22V4q0-.825.588-1.413T4 2h10.1q-.1.5-.1 1t.1 1H4v13.125L5.15 16H20V7.9q.575-.125 1.075-.338T22 7v9q0 .825-.588 1.413T20 18H6l-4 4ZM4 4v12V4Zm15 2q-1.25 0-2.125-.875T16 3q0-1.25.875-2.125T19 0q1.25 0 2.125.875T22 3q0 1.25-.875 2.125T19 6Z"/>`),
		g.Group(children),
	)
}