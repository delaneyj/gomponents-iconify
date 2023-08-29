package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M18 16V4H8a2 2 0 0 0-2 2v12"/><path d="M18 20H8a2 2 0 1 1 0-4h10c-.673 1.613-.66 2.488 0 4z"/></g>`),
		g.Group(children),
	)
}