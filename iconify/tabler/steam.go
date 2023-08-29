package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Steam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-8 8a1 1 0 1 0 2 0a1 1 0 1 0-2 0m16 0a1 1 0 1 0 2 0a1 1 0 1 0-2 0m-8 8a1 1 0 1 0 2 0a1 1 0 1 0-2 0M5.5 5.5l3 3m7 7l3 3m0-13l-3 3m-7 7l-3 3"/>`),
		g.Group(children),
	)
}