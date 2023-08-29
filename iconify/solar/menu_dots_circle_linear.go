package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuDotsCircleLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.009m3.995 0h.01m3.977 0H16"/><circle cx="12" cy="12" r="10" stroke-width="1.5"/></g>`),
		g.Group(children),
	)
}