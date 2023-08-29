package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldMore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m98 60l-68 68L0 98L98 0l98 98l-30 30zm0 264l68-68l30 30l-98 98l-98-98l30-30z"/>`),
		g.Group(children),
	)
}