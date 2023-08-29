package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inclination(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M22 19H3.41a.6.6 0 0 1-.431-1.016L16.444 4"/><path d="M20 16c-.5-3.5-1-5-3-8"/></g>`),
		g.Group(children),
	)
}