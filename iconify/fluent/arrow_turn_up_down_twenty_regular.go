package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTurnUpDownTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.566 16.996a.5.5 0 0 0 .4-.31l1.998-5a.5.5 0 0 0-.928-.371L15.5 15.152L10.91 3.628c-.329-.828-1.495-.843-1.847-.024l-5.023 11.7a.5.5 0 0 0 .919.394l5.023-11.7l4.567 11.468l-3.826-1.913a.5.5 0 0 0-.447.894l4.982 2.491a.5.5 0 0 0 .307.058Z"/>`),
		g.Group(children),
	)
}