package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3h16v19.943l-8-5.714l-8 5.714V3Zm2 2v14.057l6-4.286l6 4.286V5H6Z"/>`),
		g.Group(children),
	)
}