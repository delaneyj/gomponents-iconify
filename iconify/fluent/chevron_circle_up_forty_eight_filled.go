package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUpFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M24 4c11.046 0 20 8.954 20 20s-8.954 20-20 20S4 35.046 4 24S12.954 4 24 4Zm-8.616 23.634L24 19.018l8.616 8.616a1.25 1.25 0 0 0 1.768-1.768l-9.5-9.5a1.25 1.25 0 0 0-1.768 0l-9.5 9.5a1.25 1.25 0 0 0 1.768 1.768Z"/>`),
		g.Group(children),
	)
}