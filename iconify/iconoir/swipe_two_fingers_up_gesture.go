package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwipeTwoFingersUpGesture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 12a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm0 0V5m0 0L9 7.4M6.5 5L4 7.4M17.5 12a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm0 0V5m0 0L20 7.4M17.5 5L15 7.4"/>`),
		g.Group(children),
	)
}