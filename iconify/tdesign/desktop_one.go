package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3h22v16H13v1h4v2H7v-2h4v-1H1V3Zm2 2v9h18V5H3Zm18 11H3v1h18v-1Z"/>`),
		g.Group(children),
	)
}