package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20h11v2H5m6-17h-1a2 2 0 0 0-2 2h5a2 2 0 0 0-2-2m5 3h6a2 2 0 0 1-2 2h-1a4 4 0 0 0-4 4v1a4 4 0 0 1-4 4h-1a4 4 0 0 1-4-4H4a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2m2 2H4v3h2m13-8a2 2 0 0 1 2 2h1a2.9 2.9 0 0 0-3-3a2 2 0 0 1-2-2h-1a2.9 2.9 0 0 0 3 3"/>`),
		g.Group(children),
	)
}