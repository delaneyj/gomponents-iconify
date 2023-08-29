package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuFold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2V4Zm0 5.57L5.887 12L2 14.43V9.57ZM7 11h15v2H7v-2Zm-5 7h20v2H2v-2Z"/>`),
		g.Group(children),
	)
}