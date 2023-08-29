package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Console(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4h13a3 3 0 0 1 3 3v11a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2h17a2 2 0 0 0-2-2H5M3 18a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2V8H3v10m14 0h-5v-1h5v1M6 10.5l.71-.71l4.2 4.21l-4.2 4.21L6 17.5L9.5 14L6 10.5Z"/>`),
		g.Group(children),
	)
}