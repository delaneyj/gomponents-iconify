package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m4.036 2.488l6.611 2.833L8 6.455L1.427 3.638c.148-.151.329-.273.535-.352l2.074-.798Zm1.338-.514l1.55-.596a3 3 0 0 1 2.153 0l4.962 1.908c.205.08.386.2.534.352l-2.656 1.138l-6.543-2.802Zm9.62 2.572L8.5 7.329v7.45a2.99 2.99 0 0 0 .577-.158l4.962-1.909a1.5 1.5 0 0 0 .961-1.4V4.686a1.3 1.3 0 0 0-.007-.14ZM7.5 14.779v-7.45L1.007 4.546a1.505 1.505 0 0 0-.007.14v6.626a1.5 1.5 0 0 0 .962 1.4l4.961 1.909c.188.072.381.125.577.158Z"/>`),
		g.Group(children),
	)
}