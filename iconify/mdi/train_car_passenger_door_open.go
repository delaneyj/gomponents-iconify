package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarPassengerDoorOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 7H3c-1.1 0-2 .9-2 2v8h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1V9a2 2 0 0 0-2-2M8 12H3V9h5v3m6 4h-4V9h4v7m7-4h-5V9h5v3Z"/>`),
		g.Group(children),
	)
}