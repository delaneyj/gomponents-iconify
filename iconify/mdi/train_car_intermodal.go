package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrainCarIntermodal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 15v-4H3v4H1v2h1a2 2 0 1 0 4 0h12a2 2 0 1 0 4 0h1v-2h-2m-5-1H8v-1h8v1m5-9H3v5h18V5m-5 3H8V7h8v1Z"/>`),
		g.Group(children),
	)
}