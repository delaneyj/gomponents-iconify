package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParkingGarageEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7.25 7.44a2.35 2.35 0 0 1-1.53.44H4.45V10H3.19V4H5.8a2.13 2.13 0 0 1 1.44.46c.385.372.583.897.54 1.43a1.84 1.84 0 0 1-.53 1.55zm-1-2.16a1 1 0 0 0-.68-.2H4.45v1.76H5.6a1 1 0 0 0 .68-.22a.87.87 0 0 0 .24-.68a.82.82 0 0 0-.24-.66h-.03zm4.16-2a.5.5 0 0 0-.19-.68L5.72.1a.5.5 0 0 0-.49 0L.73 2.6a.5.5 0 0 0 .49.87l4.28-2.4l4.26 2.37a.5.5 0 0 0 .679-.198l.001-.002l-.03.04z" fill="currentColor"/>`),
		g.Group(children),
	)
}