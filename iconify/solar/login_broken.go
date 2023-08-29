package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoginBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M20 12a8 8 0 0 0-8-8m0 16a7.985 7.985 0 0 0 6.245-3"/><path stroke-linejoin="round" d="M4 12h10m0 0l-3-3m3 3l-3 3"/></g>`),
		g.Group(children),
	)
}