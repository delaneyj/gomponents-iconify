package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brochure(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaBrochure0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaBrochure1" d="M 16.110084,16.110084 H 160.04176 L 96.072129,64.087313"/><path id="galaBrochure2" d="m 16.110084,16.110084 -2e-6,175.916496 h 79.962047"/><path id="galaBrochure3" d="M 96.072132,64.087313 H 240.00381 V 240.0038 H 96.072129 l 3e-6,-175.916487"/><path id="galaBrochure4" d="M 160.04176,16.110084 V 64.087313"/></g>`),
		g.Group(children),
	)
}