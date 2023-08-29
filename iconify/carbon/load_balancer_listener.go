package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadBalancerListener(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15 8h2v2h-2zm4 0h2v2h-2zm-8 0h2v2h-2zm14 8h-8v-3h-2v3H7a2.002 2.002 0 0 0-2 2v6h2v-6h8v6h2v-6h8v6h2v-6a2.002 2.002 0 0 0-2-2zM4 26h4v4H4zm10 0h4v4h-4zm10 0h4v4h-4zM11 3h10v2H11z"/>`),
		g.Group(children),
	)
}