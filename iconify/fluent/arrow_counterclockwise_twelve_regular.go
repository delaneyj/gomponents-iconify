package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCounterclockwiseTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 2.5a.5.5 0 0 1 1 0v.854a4 4 0 1 1-.963 3.19c-.04-.298.207-.544.508-.544c.251 0 .451.205.489.453A3 3 0 1 0 3.764 4H4.5a.5.5 0 0 1 0 1h-2a.5.5 0 0 1-.5-.5v-2Z"/>`),
		g.Group(children),
	)
}