package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopologyStarRingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 19a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm8-14a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-8 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-4 7a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm12 7a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-4-7a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm8 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM6 12h4m4 0h4m-3-5l-2 3M9 7l2 3m0 4l-2 3m4-3l2 3M10 5h4m-4 14h4m3-2l2-3m0-4l-2-3M7 7l-2 3m0 4l2 3"/>`),
		g.Group(children),
	)
}