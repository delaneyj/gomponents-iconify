package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 6a7.5 7.5 0 0 1 7.5 7.5a7.5 7.5 0 0 1-7.5 7.5A7.5 7.5 0 0 1 4 13.5A7.5 7.5 0 0 1 11.5 6m0 1A6.5 6.5 0 0 0 5 13.5a6.5 6.5 0 0 0 6.5 6.5a6.5 6.5 0 0 0 6.5-6.5A6.5 6.5 0 0 0 11.5 7m3.75-1.75l.64-.75l3.83 3.2l-.64.76l-3.83-3.21m-7.5 0L3.92 8.46l-.64-.76l3.83-3.2l.64.75M11 11h1v2h2v1h-2v2h-1v-2H9v-1h2v-2Z"/>`),
		g.Group(children),
	)
}