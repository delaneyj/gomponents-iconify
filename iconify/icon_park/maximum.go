package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maximum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 4V44H44"/><path d="M10 38C10 38 15.3125 4 27 4C38.6875 4 44 38 44 38"/><path stroke-dasharray="2 6" d="M10 4L44 4"/></g>`),
		g.Group(children),
	)
}