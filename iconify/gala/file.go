package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func File(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaFile0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="15.992"><path id="galaFile1" stroke-linejoin="miter" d="M 32,48 V 207.9236"/><path id="galaFile2" stroke-linejoin="round" d="M 224,96 V 208"/><path id="galaFile3" stroke-linejoin="round" d="m 64,16 h 80"/><path id="galaFile4" stroke-linejoin="miter" d="M 64,240 H 192"/><path id="galaFile5" stroke-linejoin="round" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFile6" stroke-linejoin="round" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFile7" stroke-linejoin="round" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFile8" stroke-linejoin="round" d="M 223.91257,96.071779 144,16"/><path id="galaFile9" stroke-linejoin="round" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFilea" stroke-linejoin="round" d="M 144,64 V 16"/><path id="galaFileb" stroke-linejoin="round" d="m 176,96 h 48"/></g>`),
		g.Group(children),
	)
}