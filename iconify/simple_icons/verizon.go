package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Verizon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.302 0H22v.003L10.674 24H7.662L2 12h3.727l3.449 7.337z"/>`),
		g.Group(children),
	)
}