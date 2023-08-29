package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Typography(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 3h15v8H2V3Zm2 2v4h11V5H4Zm-2 9h20v2H2v-2Zm0 5h20v2H2v-2Z"/>`),
		g.Group(children),
	)
}