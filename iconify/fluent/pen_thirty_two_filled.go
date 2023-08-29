package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M21.65 3.434a4.889 4.889 0 1 1 6.915 6.914l-.901.901l.275.276a3.5 3.5 0 0 1 0 4.95l-2.232 2.232a1 1 0 0 1-1.414-1.414l2.232-2.232a1.5 1.5 0 0 0 0-2.122l-.276-.276l-14.978 14.98a3.7 3.7 0 0 1-1.841 1.002l-6.22 1.333a1 1 0 0 1-1.188-1.188l1.333-6.22a3.7 3.7 0 0 1 1.002-1.841L21.65 3.434Z"/>`),
		g.Group(children),
	)
}