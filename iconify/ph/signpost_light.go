package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignpostLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m244.46 108l-36-40a6 6 0 0 0-4.46-2h-70V32a6 6 0 0 0-12 0v34H40a14 14 0 0 0-14 14v64a14 14 0 0 0 14 14h82v66a6 6 0 0 0 12 0v-66h70a6 6 0 0 0 4.46-2l36-40a6 6 0 0 0 0-8Zm-43.13 38H40a2 2 0 0 1-2-2V80a2 2 0 0 1 2-2h161.33l30.6 34Z"/>`),
		g.Group(children),
	)
}