package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 3H5V1.8c0-.44.36-.8.8-.8h4.4c.44 0 .8.36.8.8V3h1.5A1.5 1.5 0 0 1 14 4.5V5h8v15h-8v.5a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 2 20.5v-16A1.5 1.5 0 0 1 3.5 3M18 7v2h2V7h-2m-4 0v2h2V7h-2m-4 0v2h2V7h-2m4 9v2h2v-2h-2m4 0v2h2v-2h-2m-8 0v2h2v-2h-2Z"/>`),
		g.Group(children),
	)
}