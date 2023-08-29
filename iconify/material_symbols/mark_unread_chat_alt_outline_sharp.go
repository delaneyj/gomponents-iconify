package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkUnreadChatAltOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14h8v-2H6v2Zm0-3h12V9H6v2Zm0-3h12v-.1q-.925-.2-1.688-.688T15 6H6v2ZM2 22V2h12.1q-.1.5-.1 1t.1 1H4v12h16V7.9q.575-.125 1.075-.338T22 7v11H6l-4 4ZM4 4v12V4Zm15 2q-1.25 0-2.125-.875T16 3q0-1.25.875-2.125T19 0q1.25 0 2.125.875T22 3q0 1.25-.875 2.125T19 6Z"/>`),
		g.Group(children),
	)
}