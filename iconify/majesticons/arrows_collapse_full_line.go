package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsCollapseFullLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 15h-4m0 0v4m0-4l4 4M5 9h4m0 0V5m0 4L5 5m14 4h-4m0 0V5m0 4l4-4M5 15h4m0 0v4m0-4l-4 4"/>`),
		g.Group(children),
	)
}