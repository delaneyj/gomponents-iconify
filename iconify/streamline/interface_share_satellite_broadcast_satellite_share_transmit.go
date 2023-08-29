package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceShareSatelliteBroadcastSatelliteShareTransmit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 9A6 6 0 0 1 5 .5ZM9.26 4.74L12 2"/><path d="M3.96 7.57L.5 13.5H7L5.92 9.73"/></g>`),
		g.Group(children),
	)
}