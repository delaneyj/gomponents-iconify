package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineMobiledataOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16 6.82l1.59 1.59L19 7l-4-4l-4 4l1.41 1.41L14 6.82v4.35l2 2zM1.39 4.22L8 10.83v6.35l-1.59-1.59L5 17l4 4l4-4l-1.41-1.41L10 17.18v-4.35l9.78 9.78l1.41-1.42L2.81 2.81z"/>`),
		g.Group(children),
	)
}