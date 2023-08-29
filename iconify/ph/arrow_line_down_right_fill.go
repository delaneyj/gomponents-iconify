package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineDownRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 40a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8Zm-28.94 56.61a8 8 0 0 0-8.72 1.73L144 140.69L85.66 82.34a8 8 0 0 0-11.32 11.32L132.69 152l-42.35 42.34A8 8 0 0 0 96 208h96a8 8 0 0 0 8-8v-96a8 8 0 0 0-4.94-7.39Z"/>`),
		g.Group(children),
	)
}