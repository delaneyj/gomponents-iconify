package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Switch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 19H5.99998"/><path d="M30 7L42 19"/><path d="M6.79897 29H42.799"/><path d="M6.79895 29L18.799 41"/></g>`),
		g.Group(children),
	)
}