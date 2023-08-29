package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 4h13a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-3.5l-3 3l-3-3H5a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3m0 1a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h3.91l2.59 2.59L14.09 18H18a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H5m0 3h13v1H5V8m0 3h12v1H5v-1m0 3h8v1H5v-1Z"/>`),
		g.Group(children),
	)
}