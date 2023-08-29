package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 14h-2V8h2m0 10h-2v-2h2m3.7-12H15V2H9v2H7.3C6.6 4 6 4.6 6 5.3v15.3c0 .8.6 1.4 1.3 1.4h9.3c.7 0 1.3-.6 1.3-1.3V5.3c.1-.7-.5-1.3-1.2-1.3Z"/>`),
		g.Group(children),
	)
}