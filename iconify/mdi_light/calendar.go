package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2h1a1 1 0 0 1 1 1v1h5V3a1 1 0 0 1 1-1h1a1 1 0 0 1 1 1v1a3 3 0 0 1 3 3v11a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3V3a1 1 0 0 1 1-1m8 2h1V3h-1v1M8 4V3H7v1h1M6 5a2 2 0 0 0-2 2v1h15V7a2 2 0 0 0-2-2H6M4 18a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V9H4v9m8-5h5v5h-5v-5m1 1v3h3v-3h-3Z"/>`),
		g.Group(children),
	)
}