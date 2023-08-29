package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InformationOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22.1 21.5L2.4 1.7L1.1 3l3 3C2.8 7.6 2 9.7 2 12c0 5.5 4.5 10 10 10c2.3 0 4.4-.8 6-2.1l2.8 2.8l1.3-1.2M13 17h-2v-4.1l2 2V17m-2-9.2L6.7 3.5C8.3 2.6 10.1 2 12 2c5.5 0 10 4.5 10 10c0 1.9-.6 3.7-1.5 5.3L12.2 9h.8V7h-2v.8Z"/>`),
		g.Group(children),
	)
}