package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarPassengerVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7H3c-1.1 0-2 .9-2 2v8h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1V9a2 2 0 0 0-2-2m-10 5H3V9h8v3m10 0h-8V9h8v3Z"/>`),
		g.Group(children),
	)
}