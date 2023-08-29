package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpaceflightSimulator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.433 10.227l3.17 3.17l-25.358 25.356l-3.17-3.17zm-.001-.001l4.392-1.9a.566.566 0 0 1 .74.753l-1.962 4.317M24.639 27.36l3.169 3.17l-12.679 12.678l-3.169-3.17zm0 0l4.391-1.901a.566.566 0 0 1 .74.754l-1.962 4.316"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.3 21.02l3.169 3.17L8.79 36.867l-3.169-3.17zm0 0l4.39-1.9a.566.566 0 0 1 .741.753L21.47 24.19M6.74 34.818l-3.123 1.757l2.297 2.297l1.757-3.123m5.408 5.408l-3.123 1.757l2.298 2.298l1.757-3.123m-2.724-5.294l-2.988 1.893l1.842 1.841l1.893-2.987"/>`),
		g.Group(children),
	)
}