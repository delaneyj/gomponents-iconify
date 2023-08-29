package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lapce(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.802 1.267L1.608 0v24L8 20.31v-2.535L3.802 20.2Zm4.208 13.9V6.231L18.003 12l-7.798 4.503v2.533L22.392 12L5.806 2.424V16.44Zm5.598-3.231L10.205 9.97v3.93Z"/>`),
		g.Group(children),
	)
}