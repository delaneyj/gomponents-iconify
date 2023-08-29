package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuUnfold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2V4Zm20 5.57v4.86L18.113 12L22 9.57ZM2 13v-2h15v2H2Zm0 7v-2h20v2H2Z"/>`),
		g.Group(children),
	)
}