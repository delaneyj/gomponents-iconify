package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 16h16V0h-1v2.6L11 6V0h-1v6.4l-4-.9V0H5v5.7L1 8.6V0H0zm5-2H1v-1.7l4-2.9V14zm5 0H6V8.7l.1-.1l3.9.9V14zm5 0h-4V9.7h.1L15 6.5V14z"/>`),
		g.Group(children),
	)
}