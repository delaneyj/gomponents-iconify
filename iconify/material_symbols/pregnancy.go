package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pregnancy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 22v-5H8v-7q0-1.25.875-2.125T11 7q1.25 0 2.125.875T14 10q.9.375 1.45 1.2T16 13v4h-3v5h-3Zm1-16q-.825 0-1.413-.588T9 4q0-.825.588-1.413T11 2q.825 0 1.413.588T13 4q0 .825-.588 1.413T11 6Z"/>`),
		g.Group(children),
	)
}