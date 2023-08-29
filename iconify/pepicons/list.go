package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func List(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><circle cx="5" cy="6" r="1.5" fill="currentColor"/><circle cx="5" cy="10" r="1.5" fill="currentColor"/><circle cx="5" cy="14" r="1.5" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8.5 6h7m-7 4h7m-7 4h7"/></g>`),
		g.Group(children),
	)
}