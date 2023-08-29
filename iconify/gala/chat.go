package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaChat0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaChat1" stroke-linejoin="miter" d="M 48.004588,16.00078 H 207.99487"/><path id="galaChat2" stroke-linejoin="round" d="M 240.00422,-48.000717 A 31.999937,31.999937 0 0 1 208.00429,-16.00078" transform="scale(1 -1)"/><path id="galaChat3" stroke-linejoin="round" d="M -16.004652,-48.000717 A 31.999937,31.999937 0 0 1 -48.004588,-16.00078" transform="scale(-1)"/><path id="galaChat4" stroke-linejoin="round" d="m -16.004652,160.0005 a 31.999937,31.999937 0 0 1 -31.999936,31.99994" transform="scale(-1 1)"/><path id="galaChat5" stroke-linejoin="round" d="m 240.00422,160.0005 a 31.999937,31.999937 0 0 1 -31.99993,31.99994"/><path id="galaChat6" stroke-linejoin="round" d="M 16.004649,48.00072 V 160.00051"/><path id="galaChat7" stroke-linejoin="round" d="M 240.00422,48.00072 V 160.00051"/><path id="galaChat8" stroke-linejoin="round" d="m 48.004588,192.00044 h 31.999937 l 47.999905,47.99991 47.99991,-47.99991 h 31.99994"/></g>`),
		g.Group(children),
	)
}