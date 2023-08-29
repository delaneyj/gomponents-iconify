package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberFourLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M182 152a6 6 0 0 1-6 6h-18v50a6 6 0 0 1-12 0v-50H72a6 6 0 0 1-5.65-8l40-112a6 6 0 0 1 11.3 4L80.51 146H146V96a6 6 0 0 1 12 0v50h18a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}