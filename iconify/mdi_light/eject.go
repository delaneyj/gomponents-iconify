package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.33 15l6.17-9.25L17.67 15H5.33M5 18h13v1H5v-1m2-3.97h9l-4.5-6.78L7 14.03Z"/>`),
		g.Group(children),
	)
}