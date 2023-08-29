package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M8 6L8 42"/><path d="M24 6V10"/><path d="M24 20V28"/><path d="M24 38V42"/><path d="M40 6L40 42"/></g>`),
		g.Group(children),
	)
}