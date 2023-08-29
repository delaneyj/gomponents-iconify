package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WatermelonOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWatermelonOne0"><g fill="none"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m24 4l17 29.92S36.046 38 24 38S7 33.92 7 33.92L24 4Z"/><circle cx="24" cy="17" r="2" fill="#fff"/><circle cx="27" cy="23" r="2" fill="#fff"/><circle cx="21" cy="23" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-width="4" d="M41 39.92S36.046 44 24 44S7 39.92 7 39.92"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWatermelonOne0)"/>`),
		g.Group(children),
	)
}