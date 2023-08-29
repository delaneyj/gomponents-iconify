package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeLeftGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 12a6 6 0 1 0 12 0a6 6 0 0 0-12 0Zm0 0H2m0 0l3-3m-3 3l3 3"/>`),
		g.Group(children),
	)
}