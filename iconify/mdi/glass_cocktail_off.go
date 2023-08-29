package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassCocktailOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.33 12.67L7.66 7L6.13 5.47L2.39 1.73L1.11 3L3 4.89V5l8 8v6H6v2h12v-1.11l2.84 2.84l1.27-1.27l-8.78-8.79M13 19v-4.11L17.11 19H13M8.2 5l-2-2H21v2l-6.4 6.4L10.2 7h6.3l2-2H8.2Z"/>`),
		g.Group(children),
	)
}