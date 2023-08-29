package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h18v22H3V1Zm2 2v4h14V3H5Zm14 6H5v12h14V9ZM7 12h4v2H7v-2Zm6 0h4v2h-4v-2Zm-6 4h4v2H7v-2Zm6 0h4v2h-4v-2Z"/>`),
		g.Group(children),
	)
}