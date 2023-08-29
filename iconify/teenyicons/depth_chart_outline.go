package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DepthChartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M.5 0v14.5h14V0M.5 4.5h2v1h2v3h2v3h1v3v-2h2v-2h2v-3h1v-2h2"/>`),
		g.Group(children),
	)
}