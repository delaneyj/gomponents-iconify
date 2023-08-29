package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2v7a1 1 0 1 1-2 0V1a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H2z"/>`),
		g.Group(children),
	)
}