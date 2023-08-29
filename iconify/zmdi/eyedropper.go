package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyedropper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M378 56q6 6 6 15t-6 15l-67 67l41 41l-30 30l-30-30l-191 190H0V283L190 92l-30-30l30-30l41 41l67-67q6-6 15-6t15 6zM84 341l172-172l-41-41L43 300z"/>`),
		g.Group(children),
	)
}