package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeTwoFingersLeftGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 17.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Zm0 0H5m0 0L7.4 15M5 17.5L7.4 20M12 6.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Zm0 0H5m0 0L7.4 4M5 6.5L7.4 9"/>`),
		g.Group(children),
	)
}