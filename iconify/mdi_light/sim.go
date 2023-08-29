package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sim(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11 3l-7 7v9a3 3 0 0 0 3 3h9a3 3 0 0 0 3-3V6a3 3 0 0 0-3-3h-5m.41 1H16a2 2 0 0 1 2 2v13a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-8.59L11.41 4M8 11v4h1v-4H8m3 0v2h1v-2h-1m3 0v4h1v-4h-1m-3 4v4h1v-4h-1m-3 2v2h1v-2H8m6 0v2h1v-2h-1Z"/>`),
		g.Group(children),
	)
}