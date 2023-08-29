package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLess(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m0 333l98-98l98 98l-30 30l-68-68l-68 68zM196 51l-98 98L0 51l30-30l68 68l68-68z"/>`),
		g.Group(children),
	)
}