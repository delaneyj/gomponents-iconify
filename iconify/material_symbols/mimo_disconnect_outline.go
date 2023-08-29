package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MimoDisconnectOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.175 3.175L5 5H4v11h9.2L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4l-5.3-5.3H17l1 1v2H6v-2l1-1H4q-.825 0-1.413-.588T2 16V5q0-.925.588-1.375l.587-.45ZM20.7 17.85L18.85 16H20V5H7.85l-2-2H20q.825 0 1.413.588T22 5v11q0 .65-.363 1.15t-.937.7Zm-7.35-7.35Zm-4.75.9Z"/>`),
		g.Group(children),
	)
}