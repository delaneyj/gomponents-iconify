package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoFrameOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21v-2H1V4h22v15h-4v2H5Zm-2-4h18V6H3v11Zm2-2h14l-4.5-6l-3.5 4.5l-2.5-3L5 15Zm-2 2V6v11Z"/>`),
		g.Group(children),
	)
}