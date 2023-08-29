package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopScreenShareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.425 11.575L11.85 9H13V7l3 3l-1.575 1.575ZM20.7 17.85L18.85 16H20V5H7.85l-2-2H20q.825 0 1.413.588T22 5v11q0 .65-.363 1.137t-.937.713Zm-.2 5.45L18.2 21H1v-2h15.175l-1-1H4q-.825 0-1.413-.588T2 16V4.85L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM10 12.85V14H8v-2q0-.275.025-.525t.15-.475L4 6.825V16h9.15L10 12.85Zm3.35-2.35Zm-4.775.9Z"/>`),
		g.Group(children),
	)
}