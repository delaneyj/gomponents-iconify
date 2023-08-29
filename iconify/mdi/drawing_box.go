package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DrawingBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 18h-6v-5.79c-.66.61-1.53.99-2.5.99c-2.04 0-3.7-1.66-3.7-3.7a3.7 3.7 0 0 1 3.7-3.7c2.04 0 3.7 1.66 3.7 3.7c0 .97-.38 1.84-.99 2.5H18m1-9H5c-1.11 0-2 .89-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}