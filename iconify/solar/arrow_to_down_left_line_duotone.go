package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowToDownLeftLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m17 14.5l-5 5l-5-5"/><path d="M12 19.5v-10c0-1.667-1-5-5-5" opacity=".5"/></g>`),
		g.Group(children),
	)
}