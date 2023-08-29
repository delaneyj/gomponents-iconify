package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Histogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path d="M4 44H44"/><path stroke-linejoin="round" d="M7 44C7 44 12.3125 10 24 10C35.6875 10 41 44 41 44"/><path stroke-linejoin="round" d="M4 4V44"/></g>`),
		g.Group(children),
	)
}