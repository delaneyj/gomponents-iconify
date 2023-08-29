package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSwapTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.78 2.72a.75.75 0 1 0-1.06 1.06l3.72 3.72H5.75a.75.75 0 0 0 0 1.5h14.69l-3.72 3.72a.75.75 0 1 0 1.06 1.06l5-5a.75.75 0 0 0 0-1.06l-5-5Zm-6.5 12.56a.75.75 0 1 0-1.06-1.06l-5 5a.75.75 0 0 0 0 1.06l5 5a.75.75 0 1 0 1.06-1.06L7.56 20.5h14.69a.75.75 0 0 0 0-1.5H7.56l3.72-3.72Z"/>`),
		g.Group(children),
	)
}