package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cabinet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 4c0-1.103-.897-2-2-2H5c-1.103 0-2 .897-2 2v7h18V4zm-5 4H8V5h2v1h4V5h2v3zM5 22h14c1.103 0 2-.897 2-2v-7H3v7c0 1.103.897 2 2 2zm3-6h2v1h4v-1h2v3H8v-3z"/>`),
		g.Group(children),
	)
}