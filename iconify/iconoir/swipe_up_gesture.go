package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeUpGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 10a6 6 0 1 0 0 12a6 6 0 0 0 0-12Zm0 0V2m0 0l3 3m-3-3L9 5"/>`),
		g.Group(children),
	)
}