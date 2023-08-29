package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v15H13v2h4v2H7v-2h4v-2H1V3Zm2 2v11h18V5H3Z"/>`),
		g.Group(children),
	)
}