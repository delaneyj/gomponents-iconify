package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatPaintOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22v-6H4V7q0-1.65 1.175-2.825T8 3h12v13h-5v6H9ZM6 10h12V5h-1v4h-2V5h-1v2h-2V5H8q-.825 0-1.413.588T6 7v3Zm0 4h12v-2H6v2Zm0 0v-2v2Z"/>`),
		g.Group(children),
	)
}