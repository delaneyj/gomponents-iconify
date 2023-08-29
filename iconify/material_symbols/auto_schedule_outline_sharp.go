package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoScheduleOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 19l-1.25-2.75L8 15l2.75-1.25L12 11l1.25 2.75L16 15l-2.75 1.25L12 19Zm-9 3V4h3V2h2v2h8V2h2v2h3v18H3Zm2-2h14V10H5v10ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}