package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 19h6v2h-6z"/><path fill="currentColor" d="m25.44 14l-1.27-4.55A2.009 2.009 0 0 0 22.246 8H9.754a2.009 2.009 0 0 0-1.923 1.45L6.531 14H4v2h2v7a2.002 2.002 0 0 0 2 2v3h2v-3h12v3h2v-3a2.002 2.002 0 0 0 2-2v-7h2v-2ZM9.755 10h12.492l1.428 5H8.326ZM24 21v2H8v-2h2v-2H8v-2h16v2h-2v2Z"/>`),
		g.Group(children),
	)
}