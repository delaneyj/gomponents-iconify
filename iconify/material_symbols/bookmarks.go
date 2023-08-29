package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmarks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V7q0-.825.588-1.413T5 5h10q.825 0 1.413.588T17 7v16l-7-3l-7 3Zm16-3V3H6V1h13q.825 0 1.413.588T21 3v17h-2Z"/>`),
		g.Group(children),
	)
}