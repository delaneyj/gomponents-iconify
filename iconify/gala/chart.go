package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaChart0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><rect id="galaChart1" width="224" height="224" x="16" y="16" ry="32"/><path id="galaChart2" d="M 160.00003,192.00003 V 111.99998"/><path id="galaChart3" d="M 192.00002,192.00003 V 63.999974"/><path id="galaChart4" d="m 63.999979,192.00003 v -32"/><path id="galaChart5" d="m 95.99997,128 v 64.00003"/><path id="galaChart6" d="m 128,144 v 48.00003"/></g>`),
		g.Group(children),
	)
}