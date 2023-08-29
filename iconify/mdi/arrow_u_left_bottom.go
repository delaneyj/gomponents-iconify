package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowULeftBottom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 10.5a6.5 6.5 0 0 1-6.5 6.5H7.83l3.09 3.09L9.5 21.5L4 16l5.5-5.5l1.41 1.41L7.83 15h5.67c2.5 0 4.5-2 4.5-4.5S16 6 13.5 6H6V4h7.5a6.5 6.5 0 0 1 6.5 6.5Z"/>`),
		g.Group(children),
	)
}