package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q2.075 0 3.538-1.463T17 14v-4q0-2.075-1.463-3.538T12 5Q9.925 5 8.462 6.463T7 10v4q0 2.075 1.463 3.538T12 19Zm0-2q-1.25 0-2.125-.875T9 14v-4q0-1.25.875-2.125T12 7q1.25 0 2.125.875T15 10v4q0 1.25-.875 2.125T12 17Zm0-7q.425 0 .713-.288T13 9q0-.425-.288-.713T12 8q-.425 0-.713.288T11 9q0 .425.288.713T12 10ZM5 23q-.825 0-1.413-.588T3 21V3q0-.825.588-1.413T5 1h14q.825 0 1.413.588T21 3v18q0 .825-.588 1.413T19 23H5Z"/>`),
		g.Group(children),
	)
}