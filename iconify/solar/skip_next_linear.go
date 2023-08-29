package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNextLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16.66 9.353c1.787 1.154 1.787 4.14 0 5.294L5.87 21.614C4.135 22.736 2 21.276 2 18.968V5.033c0-2.31 2.134-3.769 3.87-2.648l10.79 6.968Z"/><path stroke-linecap="round" d="M22 5v14"/></g>`),
		g.Group(children),
	)
}