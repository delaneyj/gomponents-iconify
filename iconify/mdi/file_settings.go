package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSettings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 0C4.89 0 4 .89 4 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V6l-6-6H6m7 1.5L18.5 7H13V1.5M7 22v2h2v-2H7m4 0v2h2v-2h-2m4 0v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}