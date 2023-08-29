package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.34 0L2 2H0v4h2l1.34 2H4V0h-.66zM5 1v1c.17 0 .34.02.5.06c.86.22 1.5 1 1.5 1.94a1.987 1.987 0 0 1-2 2v1c.25 0 .48-.04.72-.09h.03C7.05 6.58 8 5.4 8 4c0-1.4-.95-2.58-2.25-2.91C5.52 1.03 5.26 1 5 1zm0 2v2c.09 0 .18-.01.25-.03c.43-.11.75-.51.75-.97a.997.997 0 0 0-1-1z"/>`),
		g.Group(children),
	)
}