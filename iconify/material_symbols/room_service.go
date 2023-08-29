package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoomService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-2h20v2H2Zm1-3v-1q0-3.2 1.963-5.65T10 6.25V6q0-.825.588-1.413T12 4q.825 0 1.413.588T14 6v.25q3.1.65 5.05 3.1T21 15v1H3Z"/>`),
		g.Group(children),
	)
}