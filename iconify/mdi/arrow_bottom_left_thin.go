package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12.07 19l-2.83-2.82L19 6.42L17.58 5l-9.76 9.76L5 11.94V19Z"/>`),
		g.Group(children),
	)
}