package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.71 13.35a1 1 0 0 0-1.42 0l-2.92 2.92v-8a4 4 0 0 0-4-4H3a1 1 0 1 0 0 2h10.37a2 2 0 0 1 2 2v8l-2.92-2.92A1 1 0 0 0 11 14.76l4.62 4.63a1.19 1.19 0 0 0 .33.22a.94.94 0 0 0 .76 0a1 1 0 0 0 .33-.22l4.63-4.63a1 1 0 0 0 .04-1.41Z"/>`),
		g.Group(children),
	)
}