package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13h5v1l-3.74 5H7v1H2v-1l3.75-5H2v-1m7-4h5v1l-3.74 5H14v1H9v-1l3.75-5H9V9m7-4h5v1l-3.74 5H21v1h-5v-1l3.75-5H16V5Z"/>`),
		g.Group(children),
	)
}