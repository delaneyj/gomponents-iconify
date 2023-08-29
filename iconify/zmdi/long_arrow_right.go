package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 171h302l-76-77l30-30l128 128l-128 128l-30-30l76-77H0v-42z"/>`),
		g.Group(children),
	)
}