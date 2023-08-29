package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HighlightOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 9l-6 6v4h4l6-6m2-2l2.503-2.503a2.828 2.828 0 1 0-4-4l-2.497 2.497M12.5 5.5l4 4m-12 4l4 4M19 15h2v2m-2 2h-6l3-3M3 3l18 18"/>`),
		g.Group(children),
	)
}