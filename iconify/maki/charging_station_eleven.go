package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChargingStationEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.5 3H9V1.5a.5.5 0 0 0-1 0V3a1 1 0 0 0 1 1v4.25a.25.25 0 1 1-.5 0V6.5A1.5 1.5 0 0 0 7 5V2a1 1 0 0 0-1-1H2a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V6a.5.5 0 0 1 .5.5v1.75a1.25 1.25 0 1 0 2.5 0V3.5a.5.5 0 0 0-.5-.5zm-6 5.75H3L4 6H1.75L4.5 2.25H5L4 5h2.25L3.5 8.75z" fill="currentColor"/>`),
		g.Group(children),
	)
}