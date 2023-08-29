package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageFlashOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17.2V4h16v6h2V4c0-1.1-.9-2-2-2H4c-1.1 0-2 .9-2 2v18l4-4h9v-2H5.2L4 17.2M22.5 16h-2.2l1.7-4h-5v6h2v5l3.5-7Z"/>`),
		g.Group(children),
	)
}