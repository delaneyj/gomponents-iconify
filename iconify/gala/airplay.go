package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaAirplay0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaAirplay1" d="M 64,192 H 47.999992 m 0,0 c -17.728,0 -32,-14.272 -32,-32 V 47.999993 c 0,-17.728 14.272,-32 32,-32 H 208.00001 c 17.728,0 32,14.272 32,32 V 160 c 0,17.728 -14.272,32 -32,32"/><path id="galaAirplay2" d="m 64,240 64,-80 64,80 z"/><path id="galaAirplay3" d="M 208,192 H 192"/></g>`),
		g.Group(children),
	)
}