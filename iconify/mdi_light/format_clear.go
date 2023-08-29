package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatClear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4h9v1h-4L8.5 17h-1L12 5H8V4M5 21v-1h8v1H5m11.79-3.5L14 14.71l.71-.71l2.79 2.79L20.29 14l.71.71l-2.79 2.79L21 20.29l-.71.71l-2.79-2.79L14.71 21l-.71-.71l2.79-2.79Z"/>`),
		g.Group(children),
	)
}