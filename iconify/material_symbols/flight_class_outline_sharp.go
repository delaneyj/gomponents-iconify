package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightClassOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13V4h6v9h-6Zm2-2h2V6h-2v5Zm4 7H8L5 8V4h2v4l2.5 8H18v2ZM8 21v-2h10v2H8Zm6-15h2h-2Z"/>`),
		g.Group(children),
	)
}