package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Search(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaSearch0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaSearch1" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="m 89.074145,145.23139 -68.17345,68.17344"/><path id="galaSearch2" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 111.27275,167.42999 43.099304,235.60344"/><path id="galaSearch3" stroke-linecap="round" stroke-linejoin="round" d="m 43.099305,235.60344 a 15.696788,15.696788 0 0 1 -22.19861,0"/><path id="galaSearch4" stroke-linecap="round" stroke-linejoin="round" d="m 20.900695,213.40483 a 15.696788,15.696788 0 0 0 0,22.19861"/><path id="galaSearch5" stroke-linecap="round" stroke-linejoin="round" d="M 240.65575,86.483932 A 70.635544,70.635544 0 0 1 170.0202,157.11948 70.635544,70.635544 0 0 1 99.384659,86.483932 70.635544,70.635544 0 0 1 170.0202,15.848389 70.635544,70.635544 0 0 1 240.65575,86.483932 Z"/><path id="galaSearch6" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 89.074145,145.23139 22.198605,22.1986"/><path id="galaSearch7" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 100.17344,156.33068 19.89988,-19.89987"/><path id="galaSearch8" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 70.126446,164.17908 22.198606,22.1986"/><path id="galaSearch9" stroke-linecap="round" stroke-linejoin="round" d="M 209.26216,86.483936 A 39.241967,39.241967 0 0 1 170.0202,125.7259"/></g>`),
		g.Group(children),
	)
}