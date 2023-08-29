package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Min(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 4V44H44"/><path d="M10 4C10 4 15.3125 38 27 38C38.6875 38 44 4 44 4"/><path stroke-dasharray="2 6" d="M10 38L44 38"/></g>`),
		g.Group(children),
	)
}