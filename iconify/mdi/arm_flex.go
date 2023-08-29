package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArmFlex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18.34S4 7.09 7 3l5 1l-1 3.09H9v7.16h1c2-3.07 6.14-4.19 8.64-3.07c3.3 1.53 3 6.14 0 8.18C16.24 21 9 22.43 3 18.34Z"/>`),
		g.Group(children),
	)
}