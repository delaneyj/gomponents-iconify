package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileScript(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaFileScript0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-miterlimit="4" stroke-opacity="1"><path id="galaFileScript1" stroke-linejoin="miter" stroke-width="15.992" d="M 32,48 V 207.9236"/><path id="galaFileScript2" stroke-linejoin="round" stroke-width="15.992" d="M 224,96 V 208"/><path id="galaFileScript3" stroke-linejoin="round" stroke-width="15.992" d="m 64,16 h 80"/><path id="galaFileScript4" stroke-linejoin="miter" stroke-width="15.992" d="M 64,240 H 192"/><path id="galaFileScript5" stroke-linejoin="round" stroke-width="15.992" d="m 224,208 c 0.0874,15.98169 -16,32 -32,32"/><path id="galaFileScript6" stroke-linejoin="round" stroke-width="15.992" d="m -32,208 c -10e-7,16 -16,32 -32,32" transform="scale(-1 1)"/><path id="galaFileScript7" stroke-linejoin="round" stroke-width="15.992" d="M -32,-47.976784 C -32,-32 -48,-16.356322 -63.999997,-16.000002" transform="scale(-1)"/><path id="galaFileScript8" stroke-linejoin="round" stroke-width="15.992" d="M 223.91257,96.071779 144,16"/><path id="galaFileScript9" stroke-linejoin="round" stroke-width="15.992" d="m -144,64 c -0.0492,15.912926 -16.06452,31.999995 -32,32" transform="scale(-1 1)"/><path id="galaFileScripta" stroke-linejoin="round" stroke-width="15.992" d="M 144,64 V 16"/><path id="galaFileScriptb" stroke-linejoin="round" stroke-width="15.992" d="m 176,96 h 48"/><path id="galaFileScriptc" stroke-linejoin="round" stroke-width="16" d="M 64,208 96.000003,176 64,144"/><path id="galaFileScriptd" stroke-linejoin="round" stroke-width="16" d="m 112,208 h 32"/></g>`),
		g.Group(children),
	)
}