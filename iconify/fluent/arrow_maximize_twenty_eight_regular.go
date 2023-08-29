package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowMaximizeTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16.25 3h8a.75.75 0 0 1 .743.648L25 3.75v8a.75.75 0 0 1-1.493.102l-.007-.102V5.56L5.56 23.5h6.19a.75.75 0 0 1 .743.648l.007.102a.75.75 0 0 1-.648.743L11.75 25h-8a.75.75 0 0 1-.743-.648L3 24.25v-8a.75.75 0 0 1 1.493-.102l.007.102v6.188L22.438 4.5H16.25a.75.75 0 0 1-.743-.648L15.5 3.75a.75.75 0 0 1 .648-.743L16.25 3h8h-8Z"/>`),
		g.Group(children),
	)
}