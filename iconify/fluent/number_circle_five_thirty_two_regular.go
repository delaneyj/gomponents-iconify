package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13.498 9a1 1 0 0 0-.994.9l-.5 5a1 1 0 0 0 1.008 1.1l.127-.002a133.97 133.97 0 0 1 1.429-.008c.407 0 .826.003 1.175.01c.375.008.603.02.671.03a3 3 0 1 1-3.016 4.47a1 1 0 1 0-1.732 1a5 5 0 1 0 5.026-7.452a9.24 9.24 0 0 0-.908-.048a62.741 62.741 0 0 0-1.68-.009l.3-2.991h4.593a1 1 0 1 0 0-2H13.5ZM16 2C8.268 2 2 8.268 2 16s6.268 14 14 14s14-6.268 14-14S23.732 2 16 2ZM4 16C4 9.373 9.373 4 16 4s12 5.373 12 12s-5.373 12-12 12S4 22.627 4 16Z"/>`),
		g.Group(children),
	)
}