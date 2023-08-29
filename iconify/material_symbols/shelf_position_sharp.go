package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShelfPositionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16v5h19v-5H3Zm14-2h5V3h-5v11ZM3 14h5V3H3v11Zm7 0h5V3h-5v11Z"/>`),
		g.Group(children),
	)
}