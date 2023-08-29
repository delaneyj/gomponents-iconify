package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SatelliteAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.74 23.53l-3.78-3.77a.976.976 0 0 1-.95.75a1.003 1.003 0 0 1-1-1a.987.987 0 0 1 .75-.95l-3.51-3.51a6 6 0 0 1 8.49 8.48Z"/><path fill="currentColor" d="m22.46 12.45l-2.94-2.94l.86-.86a1.956 1.956 0 0 0 0-2.76L18.9 4.41a1.968 1.968 0 0 0-2.77 0l-.85.85l-2.95-2.94a1.672 1.672 0 0 0-2.36 0L6.2 6.09a1.672 1.672 0 0 0 0 2.36l2.95 2.95l-.15.14a1.95 1.95 0 0 0 0 2.77l.74.74l.74.74a1.955 1.955 0 0 0 2.76 0l.15-.15l2.94 2.94a1.678 1.678 0 0 0 2.37 0l3.76-3.76a1.666 1.666 0 0 0 0-2.37ZM7.38 7.27l3.77-3.77l2.94 2.95l-3.76 3.76ZM17.52 17.4l-2.95-2.94l3.77-3.77l2.94 2.94Z"/>`),
		g.Group(children),
	)
}