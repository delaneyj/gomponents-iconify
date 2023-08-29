package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrderAdjustmentColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3v8h5.586l-2-2L18 7.586L22.414 12L18 16.414L16.586 15l2-2H13v8h-2v-8H5.414l2 2L6 16.414L1.586 12L6 7.586L7.414 9l-2 2H11V3h2Z"/>`),
		g.Group(children),
	)
}