package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilletThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2 20v-4C2 8.268 8.268 2 16 2h4m.839 18.84h-3.536m3.536 0v-3.537m0 3.536L18 18"/><path stroke-dasharray="2 3" d="m9 9l7 7"/></g>`),
		g.Group(children),
	)
}