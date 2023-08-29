package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartTimeline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 13v-1H1V0H0v13h5v2H0v1h16v-1h-5v-2h5z"/><path fill="currentColor" d="M9 7L6 4L2 8v3h14V0L9 7z"/>`),
		g.Group(children),
	)
}