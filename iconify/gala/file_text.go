package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaFileText0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaFileText1" d="M 32,48 V 207.9236"/><path id="galaFileText2" d="M 224,96 V 208"/><path id="galaFileText3" d="m 64,16 h 80"/><path id="galaFileText4" d="M 64,240 H 192"/><path id="galaFileText5" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileText6" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileText7" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileText8" d="M 223.91257,96.071779 144,16"/><path id="galaFileText9" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileTexta" d="M 144,64 V 16"/><path id="galaFileTextb" d="m 176,96 h 48"/><path id="galaFileTextc" d="m 80,208 h 32"/><path id="galaFileTextd" d="m 96,144 v 64"/><path id="galaFileTexte" d="m 64,144 h 64"/><path id="galaFileTextf" d="m 128,144 v 16"/><path id="galaFileTextg" d="m 64,144 v 16"/></g>`),
		g.Group(children),
	)
}