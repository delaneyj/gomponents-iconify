package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2M7 7h2v2H7V7m0 4h2v2H7v-2m0 4h2v2H7v-2m10 2h-6v-2h6v2m0-4h-6v-2h6v2m0-4h-6V7h6v2Z"/>`),
		g.Group(children),
	)
}