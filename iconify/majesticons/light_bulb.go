package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightBulb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M13 3a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0V3zM6.207 4.793a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414-1.414l-1-1zm13 1.414a1 1 0 0 0-1.414-1.414l-1 1a1 1 0 0 0 1.414 1.414l1-1zM12 6a6 6 0 0 0-3.317 11h6.634A6 6 0 0 0 12 6zm3 12H9v1a3 3 0 1 0 6 0v-1zM3 11a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2H3zm17 0a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2h-1z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}