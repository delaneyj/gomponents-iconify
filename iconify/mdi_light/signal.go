package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21v-5h4v5H2m1-4v3h2v-3H3m4 4v-9h4v9H7m1-8v7h2v-7H8m4 8V8h4v13h-4m1-12v11h2V9h-2m4 12V4h4v17h-4m1-16v15h2V5h-2Z"/>`),
		g.Group(children),
	)
}