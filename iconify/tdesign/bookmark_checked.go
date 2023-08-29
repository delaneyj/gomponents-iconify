package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.596 2.94L16.94 8.595L13.405 5.06l1.414-1.415l2.121 2.122l4.243-4.243l1.414 1.414ZM4 3h8v2H6v14.057l6-4.286l6 4.286V10h2v12.943l-8-5.714l-8 5.714V3Z"/>`),
		g.Group(children),
	)
}