package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.93 19l2.83-2.82L5 6.42L6.42 5l9.76 9.76L19 11.94V19Z"/>`),
		g.Group(children),
	)
}