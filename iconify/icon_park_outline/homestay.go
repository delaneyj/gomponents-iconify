package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Homestay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M8 26c2.319.197 10 2 10 5s-4.135 1.989-6 2c-1.601-.136-6 0-6 3s7 5 14 6s18 1 18 1M8 20l6-6m14-8h10l4 4m-12 4l6 6H20l-6-6h16Zm12 8v-6M26 30v-4m6 8v-8"/>`),
		g.Group(children),
	)
}