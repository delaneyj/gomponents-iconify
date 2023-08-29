package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkingToolThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.018 2a2 2 0 0 0-2 2v.812A.974.974 0 0 0 2 5v2a3 3 0 0 0 3 3h.504l.003.005h20.986l.003-.005H27a3 3 0 0 0 3-3V4a2 2 0 0 0-2-2H4.018Zm9 24c.021.82.216 1.716.622 2.454c.429.78 1.195 1.535 2.344 1.55h.032c1.148-.015 1.915-.77 2.344-1.55c.406-.738.6-1.634.622-2.454h-5.964Zm6.942-3.096l5.52-10.899H6.52l5.514 10.898A2 2 0 0 0 13.819 24h4.358a2 2 0 0 0 1.784-1.096Z"/>`),
		g.Group(children),
	)
}