package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InfraredThermometerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2v9h-3.001L18 12a4 4 0 0 1-4 4h-1.379l-.613 3.111l.911 1.321A1 1 0 0 1 12.096 22H3l2.313-10.024L3 11l4-9h14Zm-5.001 9h-2.394l-.591 3H14a2 2 0 0 0 2-2l-.001-1Z"/>`),
		g.Group(children),
	)
}