package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.403 5.99l8.193 8.192V7.818h2v9.778H7.818v-2h6.364L5.989 7.404L7.403 5.99Z"/>`),
		g.Group(children),
	)
}