package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 3h13a3 3 0 0 1 3 3v13a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2v3h5V4H5M3 19a2 2 0 0 0 2 2h3v-5H3v3m5-9H3v5h5v-5m10 11a2 2 0 0 0 2-2v-3h-5v5h3m2-11h-5v5h5v-5m0-4a2 2 0 0 0-2-2h-3v5h5V6M9 4v5h5V4H9m0 17h5v-5H9v5m5-11H9v5h5v-5Z"/>`),
		g.Group(children),
	)
}