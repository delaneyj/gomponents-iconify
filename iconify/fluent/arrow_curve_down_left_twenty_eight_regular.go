package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCurveDownLeftTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19.401 3.378a.75.75 0 0 0-1.023-.28C13.072 6.132 13 11.269 13 14.75v7.69l-4.72-4.72a.75.75 0 1 0-1.06 1.06l6 6a.75.75 0 0 0 1.06 0l6-6a.75.75 0 0 0-1.06-1.06l-4.72 4.72v-7.69c0-3.518.128-7.78 4.622-10.349a.75.75 0 0 0 .28-1.023Z"/>`),
		g.Group(children),
	)
}