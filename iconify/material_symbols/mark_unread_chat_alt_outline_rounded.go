package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkUnreadChatAltOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 14h6q.425 0 .713-.288T14 13q0-.425-.288-.713T13 12H7q-.425 0-.713.288T6 13q0 .425.288.713T7 14Zm0-3h10q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9H7q-.425 0-.713.288T6 10q0 .425.288.713T7 11Zm0-3h11v-.1q-.925-.2-1.688-.688T15 6H7q-.425 0-.713.288T6 7q0 .425.288.713T7 8ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h10.1q-.1.5-.1 1t.1 1H4v12h16V7.9q.575-.125 1.075-.338T22 7v9q0 .825-.588 1.413T20 18H6ZM4 4v12V4Zm15 2q-1.25 0-2.125-.875T16 3q0-1.25.875-2.125T19 0q1.25 0 2.125.875T22 3q0 1.25-.875 2.125T19 6Z"/>`),
		g.Group(children),
	)
}