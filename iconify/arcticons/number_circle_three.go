package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M25.976 24.003a6.074 6.074 0 0 0 6.073-6.074h0a6.074 6.074 0 0 0-6.073-6.073m0 24.294a6.074 6.074 0 0 0 6.073-6.074h0a6.074 6.074 0 0 0-6.073-6.073M15.953 34.1c1.678 1.406 3.489 2.05 7.556 2.05h2.467"/><path d="M15.95 13.88c1.682-1.402 3.495-2.04 7.562-2.03l2.466.006m-6.19 12.147h6.188"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}