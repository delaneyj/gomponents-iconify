package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 .5h8V3h5v20.5H3V3h5V.5Zm2 2V5H5v16.5h14V5h-5V2.5h-4Zm3 7v3h3v2h-3v3h-2v-3H8v-2h3v-3h2Z"/>`),
		g.Group(children),
	)
}