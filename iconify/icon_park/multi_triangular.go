package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiTriangular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 6L4 41H44L24 6Z"/><path d="M39 32.25L34 41"/><path d="M29 14.75L14 41"/><path d="M34 23.5L24 41"/></g>`),
		g.Group(children),
	)
}