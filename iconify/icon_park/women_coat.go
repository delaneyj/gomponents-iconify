package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WomenCoat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M12 10C16 9 19 4 19 4H29C29 4 32 9.2 36 10L42 31H34V44H14V31H6L12 10Z"/><path d="M19 4L24 18L29 4"/><path d="M24 18L24 44"/></g>`),
		g.Group(children),
	)
}