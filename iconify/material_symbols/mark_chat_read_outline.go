package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkChatReadOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.35 20l-3.525-3.55l1.4-1.4l2.125 2.125l4.25-4.25L23 14.35L17.35 20ZM2 22V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v7h-2V4H4v13.125L5.15 16H12v2H6l-4 4Zm2-6V4v12Z"/>`),
		g.Group(children),
	)
}