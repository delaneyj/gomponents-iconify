package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeafTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M11.037 3.048a1.75 1.75 0 0 1 2.2.225l3.712 3.711a6.999 6.999 0 0 1-4.2 11.908v2.358a.75.75 0 0 1-1.5 0v-2.358A6.999 6.999 0 0 1 7.05 6.985l3.712-3.712a1.76 1.76 0 0 1 .275-.225zM12 11a.75.75 0 0 0-.75.75v5.63a5.532 5.532 0 0 0 1.5 0v-5.63A.75.75 0 0 0 12 11z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}