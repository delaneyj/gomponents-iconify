package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarHopperFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 8c-1.96-1.81-5.26-3-9-3S4.96 6.19 3 8H1v9h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1V8h-2m-8 7v-4h-2v4H8v-4H6v4H3v-5h18v5h-3v-4h-2v4h-3Z"/>`),
		g.Group(children),
	)
}