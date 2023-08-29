package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SleepOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.79 4.46l.71-.71L20.25 20.5l-.71.71L14 15.66V16H9v-1l1.86-2.5l-8.07-8.04M2 13h5v1l-3.74 5H7v1H2v-1l3.75-5H2v-1m12-4v1l-1.21 1.62l-.72-.71l.68-.91h-1.59l-1-1H14m-3.74 6h3.08l-1.76-1.76L10.26 15M16 5h5v1l-3.74 5H21v1h-5v-1l3.75-5H16V5Z"/>`),
		g.Group(children),
	)
}