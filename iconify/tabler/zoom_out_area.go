package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOutArea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 15h4m-7 0a5 5 0 1 0 10 0a5 5 0 1 0-10 0m12 7l-3-3M6 18H5a2 2 0 0 1-2-2v-1m0-4v-1m0-4V5a2 2 0 0 1 2-2h1m4 0h1m4 0h1a2 2 0 0 1 2 2v1"/>`),
		g.Group(children),
	)
}