package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaShield0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaShield1" d="m 128,239.89376 c 0,0 -95.908938,-31.96965 -95.908938,-111.89376"/><path id="galaShield2" d="M 32.091062,128 V 32.09106 c 0,0 31.969645,-15.984825 95.908938,-15.984825"/><path id="galaShield3" d="m 128,239.89376 c 0,0 95.90894,-31.96965 95.90894,-111.89376"/><path id="galaShield4" d="M 223.90894,128 V 32.09106 c 0,0 -31.96965,-15.984823 -95.90894,-15.984823"/><path id="galaShield5" d="m 128,239.89376 2.8e-4,-223.787523"/></g>`),
		g.Group(children),
	)
}