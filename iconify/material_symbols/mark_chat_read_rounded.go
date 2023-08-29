package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkChatReadRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.35 17.175l3.55-3.55q.3-.3.7-.288t.7.313q.275.3.288.7t-.288.7l-4.25 4.25q-.3.3-.7.3t-.7-.3l-2.125-2.15q-.275-.275-.287-.687t.287-.713q.275-.275.7-.275t.7.275l1.425 1.425ZM6 18l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v6q0 .425-.288.713T21 11h-7q-.825 0-1.413.588T12 13v4q0 .425-.288.713T11 18H6Z"/>`),
		g.Group(children),
	)
}