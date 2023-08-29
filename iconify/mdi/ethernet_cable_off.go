package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EthernetCableOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3h2v4h-2V3M8 4h2v4h4V4h2v7h-3.18L8 6.18V4m12 16.72L18.73 22L14 17.27V22h-4v-8.73l-8-8L3.28 4L20 20.72Z"/>`),
		g.Group(children),
	)
}