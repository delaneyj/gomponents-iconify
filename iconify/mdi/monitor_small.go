package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2H5c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h5v2H8v2h8v-2h-2v-2h5c1.11 0 2-.89 2-2V4a2 2 0 0 0-2-2m0 14H5V4h14v12Z"/>`),
		g.Group(children),
	)
}