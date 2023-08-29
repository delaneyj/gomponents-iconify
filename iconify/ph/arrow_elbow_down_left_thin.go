package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowElbowDownLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M196 32v144a4 4 0 0 1-4 4H57.66l41.17 41.17a4 4 0 0 1-5.66 5.66l-48-48a4 4 0 0 1 0-5.66l48-48a4 4 0 1 1 5.66 5.66L57.66 172H188V32a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}