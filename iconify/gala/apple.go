package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaApple0"><g id="galaApple1" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="4.244" transform="translate(.331 .207) scale(3.76975)"><path id="galaApple2" d="m 33.866666,21.166666 c -4.233333,1e-6 -25.465872,-8.433732 -25.465872,12.732936"/><path id="galaApple3" d="m 8.400794,33.899602 c 0,21.166665 12.743914,29.710184 16.977248,29.710184 4.233334,0 4.255291,-4.244312 8.488624,-4.244312"/><path id="galaApple4" d="m 33.866665,21.166666 c 4.233334,1e-6 25.465874,-8.433732 25.465874,12.732936"/><path id="galaApple5" d="m 59.354497,33.866666 c 0,21.166666 -12.765873,29.74312 -16.999207,29.74312 -4.233334,0 -4.25529,-4.244312 -8.488624,-4.244312"/><path id="galaApple6" d="m 33.866666,21.166666 c 0,-12.6999995 -8.488624,-16.9772483 -8.488624,-16.9772483"/><path id="galaApple7" d="m 33.866666,21.166666 c 0,-16.9333325 8.499603,-12.7439153 12.732936,-16.9772483"/><path id="galaApple8" d="m 33.866666,21.166666 c 12.7,-4.233333 16.966269,-8.510583 12.732936,-16.9772483"/></g></g>`),
		g.Group(children),
	)
}