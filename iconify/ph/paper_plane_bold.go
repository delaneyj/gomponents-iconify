package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m241.42 198.2l-.06-.09L145.3 30.17a20 20 0 0 0-34.82 0L14.58 198.2a20 20 0 0 0 24.1 28.64l89.2-30.61l89.45 30.61a20.22 20.22 0 0 0 6.72 1.16a20 20 0 0 0 17.37-29.8ZM140 175v-55a12 12 0 0 0-24 0v54.93l-75.24 25.82l87.13-152.69l87.34 152.7Z"/>`),
		g.Group(children),
	)
}