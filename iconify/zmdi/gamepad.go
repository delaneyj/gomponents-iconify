package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m277 120l-64 64l-64-64V3h128v117zm-160 32l64 64l-64 64H0V152h117zm32 160l64-64l64 64v117H149V312zm160-160h118v128H309l-64-64z"/>`),
		g.Group(children),
	)
}