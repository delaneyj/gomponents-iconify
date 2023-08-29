package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 10a7.5 7.5 0 1 0-2.952 5.964l4.745 4.743l.094.083a1 1 0 0 0 1.32-1.497l-4.743-4.745A7.468 7.468 0 0 0 17.5 10Zm-4-1a1 1 0 1 1 0 2h-7a1 1 0 1 1 0-2h7Z"/>`),
		g.Group(children),
	)
}