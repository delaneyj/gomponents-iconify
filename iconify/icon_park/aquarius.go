package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aquarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 20L14 8L20 17L28 8L35 19L44 9.04348"/><path d="M4 40L14 28L20 37L28 28L35 39L44 29.0435"/></g>`),
		g.Group(children),
	)
}