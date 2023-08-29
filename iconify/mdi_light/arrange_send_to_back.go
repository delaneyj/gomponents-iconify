package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrangeSendToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 5H4v5h5V5m1 6H3V4h7v7m3 10v-7h7v7h-7m1-1h5v-5h-5v5m2-12v4h-1V9h-3V8h4m-9 9v-4h1v3h3v1H7Z"/>`),
		g.Group(children),
	)
}