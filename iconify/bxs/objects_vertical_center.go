package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectsVerticalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 7a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v4h-2V5a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v6H2v2h3v6a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-6h2v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4h3v-2h-3z"/>`),
		g.Group(children),
	)
}