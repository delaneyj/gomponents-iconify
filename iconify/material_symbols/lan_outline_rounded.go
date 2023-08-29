package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20v-3q0-.825.588-1.413T5 15h1v-2q0-.825.588-1.413T8 11h3V9h-1q-.825 0-1.413-.588T8 7V4q0-.825.588-1.413T10 2h4q.825 0 1.413.588T16 4v3q0 .825-.588 1.413T14 9h-1v2h3q.825 0 1.413.588T18 13v2h1q.825 0 1.413.588T21 17v3q0 .825-.588 1.413T19 22h-4q-.825 0-1.413-.588T13 20v-3q0-.825.588-1.413T15 15h1v-2H8v2h1q.825 0 1.413.588T11 17v3q0 .825-.588 1.413T9 22H5Zm5-15h4V4h-4v3ZM5 20h4v-3H5v3Zm10 0h4v-3h-4v3ZM12 7ZM9 17Zm6 0Z"/>`),
		g.Group(children),
	)
}