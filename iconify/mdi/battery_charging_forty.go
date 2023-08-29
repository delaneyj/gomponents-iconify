package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryChargingForty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4h-2V2H5v2H3c-.6 0-1 .4-1 1v16c0 .6.4 1 1 1h10c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1m-1 10.5H4V6h8v8.5M23 11h-3V4l-5 10h3v8"/>`),
		g.Group(children),
	)
}