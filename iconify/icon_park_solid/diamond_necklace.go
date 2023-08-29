package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondNecklace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSDiamondNecklace0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M17 32.41L24 29l7 3.41v7.5L24 44l-7-4.09v-7.5Z"/><path d="M8 4c.455 8.333 6 25 16 25S40 12.784 40 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSDiamondNecklace0)"/>`),
		g.Group(children),
	)
}