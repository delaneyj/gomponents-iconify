package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoChatOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 14h6q.425 0 .713-.288T15 13v-2l2 2V7l-2 2V7q0-.425-.288-.713T14 6H8q-.425 0-.713.288T7 7v6q0 .425.288.713T8 14Zm-6 8V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6l-4 4Zm3.15-6H20V4H4v13.125L5.15 16ZM4 16V4v12Z"/>`),
		g.Group(children),
	)
}