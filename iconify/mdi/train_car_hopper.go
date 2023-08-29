package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarHopper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 8v9h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1V8H1m12 7v-4h-2v4H8v-4H6v4H3v-5h18v5h-3v-4h-2v4h-3Z"/>`),
		g.Group(children),
	)
}