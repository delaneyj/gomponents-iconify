package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatVerticalAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2V4Zm0 7h14v2H2v-2Zm1 7H2v2h20v-2H3Z"/>`),
		g.Group(children),
	)
}