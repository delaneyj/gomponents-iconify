package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PixelThreeThreeXlThreeA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 6q.425 0 .713-.288T10 5q0-.425-.288-.713T9 4q-.425 0-.713.288T8 5q0 .425.288.713T9 6ZM7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23H7Z"/>`),
		g.Group(children),
	)
}