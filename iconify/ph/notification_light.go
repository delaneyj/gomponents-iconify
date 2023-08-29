package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotificationLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M214 128v80a14 14 0 0 1-14 14H48a14 14 0 0 1-14-14V56a14 14 0 0 1 14-14h80a6 6 0 0 1 0 12H48a2 2 0 0 0-2 2v152a2 2 0 0 0 2 2h152a2 2 0 0 0 2-2v-80a6 6 0 0 1 12 0Zm16-68a34 34 0 1 1-34-34a34 34 0 0 1 34 34Zm-12 0a22 22 0 1 0-22 22a22 22 0 0 0 22-22Z"/>`),
		g.Group(children),
	)
}