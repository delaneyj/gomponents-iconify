package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-6h3v6H5Zm6 0V9h3v11h-3Zm6 0V4h3v16h-3Z"/>`),
		g.Group(children),
	)
}