package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaveSawtoothThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m234.1 131.41l-104 64a4 4 0 0 1-2.1.59a4 4 0 0 1-4-4V71.16l-97.9 60.25a4 4 0 0 1-4.2-6.82l104-64A4 4 0 0 1 132 64v120.84l97.9-60.25a4 4 0 1 1 4.2 6.82Z"/>`),
		g.Group(children),
	)
}