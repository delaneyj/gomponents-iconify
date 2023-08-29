package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 408L235 3l234 405H0zm256-64v-43h-43v43h43zm0-85v-86h-43v86h43z"/>`),
		g.Group(children),
	)
}