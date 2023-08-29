package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14m-3.42 14l-.95-4.08l3.16-2.73l-4.17-.36L12 6l-1.62 3.84l-4.17.36l3.16 2.73L8.42 17L12 14.84L15.58 17Z"/>`),
		g.Group(children),
	)
}