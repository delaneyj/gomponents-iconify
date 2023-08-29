package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowExpandRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2H2v20h2v-9h14.17l-5.5 5.5l1.41 1.42L22 12l-7.92-7.92l-1.41 1.42l5.5 5.5H4V2Z"/>`),
		g.Group(children),
	)
}