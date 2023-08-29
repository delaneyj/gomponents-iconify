package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RugbyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-miterlimit="2" stroke-width="4"><path d="M4 24C6.2922 32.63 14.3769 39 24 39C33.6231 39 41.7078 32.63 44 24C41.7078 15.37 33.6231 9 24 9C14.3769 9 6.2922 15.37 4 24Z"/><path stroke-linecap="round" d="M40 24H44"/><path stroke-linecap="round" d="M14 24H34"/><path stroke-linecap="round" d="M4 24H8"/><path stroke-linecap="round" d="M17 20V28"/><path stroke-linecap="round" d="M31 20V28"/><path stroke-linecap="round" d="M24 20V28"/></g>`),
		g.Group(children),
	)
}