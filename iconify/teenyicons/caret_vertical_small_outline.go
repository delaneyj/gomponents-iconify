package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretVerticalSmallOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" d="m10 9l-2.5 2L5 9m0-3l2.5-2L10 6"/>`),
		g.Group(children),
	)
}