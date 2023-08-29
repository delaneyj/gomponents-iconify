package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cylinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 10c0 3.314-8.954 6-20 6S4 13.314 4 10s8.954-6 20-6s20 2.686 20 6Zm0 28c0 3.314-8.954 6-20 6S4 41.314 4 38s8.954-6 20-6s20 2.686 20 6Zm0-28v28M4 10v28"/>`),
		g.Group(children),
	)
}