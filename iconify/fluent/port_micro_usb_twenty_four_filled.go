package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortMicroUsbTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M19 12.83V13a2.006 2.006 0 0 1-2 2H7a2.006 2.006 0 0 1-2-2v-.17a2.006 2.006 0 0 1 .59-1.42l1.82-1.82A2.006 2.006 0 0 1 8.83 9h6.34a2.006 2.006 0 0 1 1.42.59l1.82 1.82a2.008 2.008 0 0 1 .59 1.42Z"/>`),
		g.Group(children),
	)
}