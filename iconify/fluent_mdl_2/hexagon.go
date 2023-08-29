package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2024 960l-497 960H521L24 960L521 0h1006l497 960zm-144 0l-431-832H599L168 960l431 832h850l431-832z"/>`),
		g.Group(children),
	)
}