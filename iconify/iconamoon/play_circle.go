package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path d="m14 12l-3 1.732v-3.464L14 12Z"/></g>`),
		g.Group(children),
	)
}