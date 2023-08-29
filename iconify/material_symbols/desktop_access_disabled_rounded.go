package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopAccessDisabledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.175 3.175v2.8L1.4 4.2q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l18.4 18.4q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L15.2 18H14v2h1q.425 0 .713.288T16 21q0 .425-.288.713T15 22H9q-.425 0-.713-.288T8 21q0-.425.288-.713T9 20h1v-2H4q-.825 0-1.413-.588T2 16V5q0-.925.588-1.375l.587-.45ZM20.7 17.85L5.85 3H20q.825 0 1.413.588T22 5v11q0 .65-.363 1.15t-.937.7Z"/>`),
		g.Group(children),
	)
}