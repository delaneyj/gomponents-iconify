package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PortHdmiTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m21.41 11.41l-2.82-2.82A2.007 2.007 0 0 0 17.17 8H6.83a2.006 2.006 0 0 0-1.42.59l-2.82 2.82A2.007 2.007 0 0 0 2 12.83V14a2.006 2.006 0 0 0 2 2h16a2.006 2.006 0 0 0 2-2v-1.17a2.006 2.006 0 0 0-.59-1.42ZM17 12a.755.755 0 0 1-.75.75h-8.5a.75.75 0 1 1 0-1.5h8.5A.756.756 0 0 1 17 12Z"/>`),
		g.Group(children),
	)
}