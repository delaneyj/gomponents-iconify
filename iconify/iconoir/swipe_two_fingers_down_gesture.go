package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeTwoFingersDownGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0 0v7m0 0L9 16.6M6.5 19L4 16.6M17.5 12a3.5 3.5 0 1 1 0-7a3.5 3.5 0 0 1 0 7Zm0 0v7m0 0l2.5-2.4M17.5 19L15 16.6"/>`),
		g.Group(children),
	)
}