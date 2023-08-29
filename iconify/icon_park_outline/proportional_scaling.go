package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProportionalScaling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M22.268 7c.77-1.333 2.694-1.333 3.464 0l17.32 30c.77 1.333-.192 3-1.731 3H6.678c-1.54 0-2.501-1.667-1.732-3l17.32-30Z"/><path stroke-linecap="round" d="m19 40l13-22m0 22l6-11"/></g>`),
		g.Group(children),
	)
}