package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drawing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 3A5.5 5.5 0 0 1 14 8.5c0 1.33-.47 2.55-1.26 3.5H21v9h-9v-8.26c-.95.79-2.17 1.26-3.5 1.26A5.5 5.5 0 0 1 3 8.5A5.5 5.5 0 0 1 8.5 3Z"/>`),
		g.Group(children),
	)
}