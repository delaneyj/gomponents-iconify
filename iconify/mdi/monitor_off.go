package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 18v2h2v2H8v-2h2v-2H3a2 2 0 0 1-2-2V4L0 3l1.41-1.42l20.75 20.76l-1.41 1.41L15 18h-1M3 16h10L3 6v10M21 2a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-.34l-2-2H21V4H6.66l-2-2H21Z"/>`),
		g.Group(children),
	)
}