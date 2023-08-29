package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiTriangularThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 41H44L24 6L4 41Z"/><path d="M24 30L24 6"/><path d="M24 30L4 41"/><path d="M24 30L44 41"/></g>`),
		g.Group(children),
	)
}