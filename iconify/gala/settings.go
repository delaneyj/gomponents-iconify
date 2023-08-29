package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Settings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaSettings0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaSettings1" d="M 48.000002,16 H 208 c 17.728,0 32,14.272 32,32 v 160 c 0,17.728 -14.272,32 -32,32 H 48.000002 c -17.728,0 -32,-14.272 -32,-32 V 48 c 0,-17.728 14.272,-32 32,-32 z"/><path id="galaSettings2" d="M 64.000006,64.000001 H 79.999993"/><path id="galaSettings3" d="m 79.999996,-96.000015 a 16,16 0 0 1 -16,16 16,16 0 0 1 -16,-16 16,16 0 0 1 16,-16.000005 16,16 0 0 1 16,16.000005 z" transform="rotate(90)"/><path id="galaSettings4" d="m 112.00001,64.000353 79.99997,-3.52e-4"/><path id="galaSettings5" d="M 191.99998,128 H 176"/><path id="galaSettings6" d="m 144,159.99997 a 16,16 0 0 1 -16,16 16,16 0 0 1 -16,-16 16,16 0 0 1 16,-16 16,16 0 0 1 16,16 z" transform="matrix(0 1 1 0 0 0)"/><path id="galaSettings7" d="M 143.99998,128.00035 64.000006,128"/><path id="galaSettings8" d="M 64.000006,192.00001 H 79.999993"/><path id="galaSettings9" d="m 208,-96.000015 a 16,16 0 0 1 -16,16 16,16 0 0 1 -16,-16 16,16 0 0 1 16,-16.000005 16,16 0 0 1 16,16.000005 z" transform="rotate(90)"/><path id="galaSettingsa" d="m 112.00001,192.00036 79.99997,-3.5e-4"/></g>`),
		g.Group(children),
	)
}