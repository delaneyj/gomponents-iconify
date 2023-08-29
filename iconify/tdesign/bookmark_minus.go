package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h9v2H6v14.057l6-4.286l6 4.286V7h2v15.943l-8-5.714l-8 5.714V3Zm11 0h8v2h-8V3Z"/>`),
		g.Group(children),
	)
}