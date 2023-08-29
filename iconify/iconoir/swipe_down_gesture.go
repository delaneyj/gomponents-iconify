package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeDownGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0 0v8m0 0l-3-3m3 3l3-3"/>`),
		g.Group(children),
	)
}