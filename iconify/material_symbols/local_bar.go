package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalBar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21v-2h5v-5L3 5V3h18v2l-8 9v5h5v2H6ZM7.45 7h9.1l1.8-2H5.65l1.8 2Z"/>`),
		g.Group(children),
	)
}