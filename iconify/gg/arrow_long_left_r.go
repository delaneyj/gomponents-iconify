package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongLeftR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m4.649 7.725l-4.25 4.236l4.235 4.25l1.417-1.412l-1.814-1.82l11.866.04l3.255 3.256l4.243-4.243L19.36 7.79l-3.23 3.23l-11.911-.04l1.843-1.837L4.65 7.726Zm13.295 4.307l1.415-1.414l1.414 1.414l-1.415 1.414l-1.414-1.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}