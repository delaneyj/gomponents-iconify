package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRightLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 10v9h2v-7h10.172l-3.95 3.95l1.414 1.414L21 11l-6.364-6.364l-1.414 1.414l3.95 3.95H5Z"/>`),
		g.Group(children),
	)
}