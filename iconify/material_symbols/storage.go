package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Storage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-4h18v4H3Zm2-1h2v-2H5v2ZM3 8V4h18v4H3Zm2-1h2V5H5v2Zm-2 7v-4h18v4H3Zm2-1h2v-2H5v2Z"/>`),
		g.Group(children),
	)
}