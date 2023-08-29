package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarMonthOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 14v-2h2v2h-2Zm-4 0v-2h2v2H7Zm8 0v-2h2v2h-2Zm-4 4v-2h2v2h-2Zm-4 0v-2h2v2H7Zm8 0v-2h2v2h-2ZM3 22V4h3V2h2v2h8V2h2v2h3v18H3Zm2-2h14V10H5v10ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}