package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkUnreadChatAltSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 6q-1.25 0-2.125-.875T16 3q0-1.25.875-2.125T19 0q1.25 0 2.125.875T22 3q0 1.25-.875 2.125T19 6ZM3.7 20.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h10.1q-.1.5-.1 1t.1 1q.125.575.35 1.075T15 6H7q-.425 0-.713.288T6 7q0 .425.288.713T7 8h12q.8 0 1.575-.25T22 7v9q0 .825-.588 1.413T20 18H6l-2.3 2.3ZM7 11h10q.425 0 .713-.288T18 10q0-.425-.288-.713T17 9H7q-.425 0-.713.288T6 10q0 .425.288.713T7 11Zm0 3h6q.425 0 .713-.288T14 13q0-.425-.288-.713T13 12H7q-.425 0-.713.288T6 13q0 .425.288.713T7 14Z"/>`),
		g.Group(children),
	)
}