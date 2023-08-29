package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeavesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M21 17c8.385-2.12 17.665-8.76 21-12c0 15-3.801 23.472-6 26c-10 11.5-20.935 6.16-23 2c-4.855-9.777 2.07-14.5 8-16ZM6 43c.412-2 2.388-6.6 7-9"/>`),
		g.Group(children),
	)
}