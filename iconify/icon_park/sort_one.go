package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 11.5H29"/><path d="M6 24.5H29"/><path d="M36 11.5V37.5L42 30.5"/><path d="M6 37.5H29"/></g>`),
		g.Group(children),
	)
}