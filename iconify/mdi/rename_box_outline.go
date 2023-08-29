package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RenameBoxOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 17h-5.5l2-2H17M7 17v-2.5l7.9-7.8c.2-.2.5-.2.7 0l1.8 1.8c.2.2.2.5 0 .7L9.5 17M19 5v14H5V5h14m0-2H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2Z"/>`),
		g.Group(children),
	)
}