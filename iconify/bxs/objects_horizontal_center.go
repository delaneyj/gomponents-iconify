package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectsHorizontalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13h-6v-2h4a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-4V2h-2v3H7a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4v2H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6v3h2v-3h6a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}