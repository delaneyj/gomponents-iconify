package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeTwoFingersRightGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0h7m0 0L16.6 15m2.4 2.5L16.6 20M12 6.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm0 0h7m0 0L16.6 4M19 6.5L16.6 9"/>`),
		g.Group(children),
	)
}