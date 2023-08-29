package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopAccessDisabled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.175 3.175v2.8L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4l-5.3-5.3H14v2h2v2H8v-2h2v-2H4q-.825 0-1.413-.588T2 16V5q0-.925.588-1.375l.587-.45ZM20.7 17.85L5.85 3H20q.825 0 1.413.588T22 5v11q0 .65-.363 1.15t-.937.7Z"/>`),
		g.Group(children),
	)
}