package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 12a9 9 0 1 0-18 0a9 9 0 0 0 18 0ZM12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1Zm.768 8.819H8.525v-2h7.657v7.657h-2v-4.243l-5.364 5.364l-1.414-1.414l5.364-5.364Z"/>`),
		g.Group(children),
	)
}