package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaVideo0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaVideo1" d="M 47.999992,48 H 160 c 17.728,0 32,14.272 32,32 v 96 c 0,17.728 -14.272,32 -32,32 H 47.999992 c -17.728,0 -32,-14.272 -32,-32 V 80 c 0,-17.728 14.272,-32 32,-32 z"/><path id="galaVideo2" d="m 192,160 48,32 V 64 l -48,32"/></g>`),
		g.Group(children),
	)
}