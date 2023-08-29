package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLongRightR(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m19.351 7.725l4.25 4.236l-4.235 4.25l-1.417-1.412l1.814-1.82l-11.866.04l-3.255 3.256l-4.243-4.243L4.642 7.79l3.229 3.23l11.911-.04l-1.843-1.837l1.412-1.417Zm-14.71 5.721l1.415-1.414l-1.414-1.414l-1.415 1.414l1.415 1.414Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}