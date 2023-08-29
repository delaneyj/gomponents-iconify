package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarHopperCovered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23 9V7H1v2l1 1.33V15H1v2h1a2 2 0 1 0 4 0h4l1 1h2l1-1h4a2 2 0 1 0 4 0h1v-2h-1v-4.67L23 9M4 15v-2l1.5 2H4m1-5V9h14v1H5m15 5h-1.5l1.5-2v2Z"/>`),
		g.Group(children),
	)
}