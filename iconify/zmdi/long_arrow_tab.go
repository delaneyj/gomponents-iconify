package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LongArrowTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m226 94l30-30l128 128l-128 128l-30-30l76-77H0v-42h302zm179-30h43v256h-43V64z"/>`),
		g.Group(children),
	)
}