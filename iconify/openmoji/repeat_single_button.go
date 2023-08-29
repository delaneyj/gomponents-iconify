package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepeatSingleButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M10.575 43.203A17.367 17.367 0 0 1 10 38.757V37.65c0-9.464 7.575-17.136 16.92-17.136h30.908m2.666 5.465A17.38 17.38 0 0 1 62 33.082v1.114c0 9.515-7.603 17.229-16.98 17.229h-16.4m29.208-30.911L49.643 12.3m0 16.427l8.185-8.213M16.647 46.472l2.979-2.231v10.76"/>`),
		g.Group(children),
	)
}