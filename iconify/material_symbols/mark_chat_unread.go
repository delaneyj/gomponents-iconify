package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkChatUnread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6q-1.25 0-2.125-.875T16 3q0-1.25.875-2.125T19 0q1.25 0 2.125.875T22 3q0 1.25-.875 2.125T19 6ZM2 22V4q0-.825.588-1.413T4 2h10.1q-.1.5-.1 1t.1 1q.35 1.725 1.725 2.863T19 8q.8 0 1.575-.25T22 7v9q0 .825-.588 1.413T20 18H6l-4 4Z"/>`),
		g.Group(children),
	)
}