package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoChatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 14h6q.425 0 .713-.288T15 13v-2l1.575 1.575q.125.125.275.063T17 12.4V7.6q0-.175-.15-.237t-.275.062L15 9V7q0-.425-.287-.713T14 6H8q-.425 0-.713.288T7 7v6q0 .425.288.713T8 14Zm-2 4l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Z"/>`),
		g.Group(children),
	)
}