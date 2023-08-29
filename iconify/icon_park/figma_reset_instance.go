package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaResetInstance(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M44 24L24 4L4 24L24 44L35 33"/><path d="M22 24.0002C22 24.0002 34.5 23.5002 39.5 28.5002C44.5 33.5002 39 44.0002 39 44.0002"/><path d="M26 20L22 24L26 28"/></g>`),
		g.Group(children),
	)
}