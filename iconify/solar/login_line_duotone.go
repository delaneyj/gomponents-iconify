package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoginLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 20a8 8 0 1 0 0-16" opacity=".5"/><path stroke-linejoin="round" d="M4 12h10m0 0l-3-3m3 3l-3 3"/></g>`),
		g.Group(children),
	)
}