package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Check(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m2.38 14.73l2.828-2.83l7.75 7.748l12.92-12.915l2.83 2.828l-15.75 15.748z"/>`),
		g.Group(children),
	)
}