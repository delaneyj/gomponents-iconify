package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Display(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaDisplay0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-miterlimit="4" stroke-width="16"><path id="galaDisplay1" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 32.045863,15.996967 H 223.95413"/><path id="galaDisplay2" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 16.053507,31.994108 2e-6,143.926412"/><path id="galaDisplay3" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 239.94649,31.989325 V 175.92052"/><path id="galaDisplay4" stroke-linecap="round" stroke-linejoin="round" d="M 32.045861,15.996971 A 15.992355,15.992355 0 0 0 16.053507,31.989325"/><path id="galaDisplay5" stroke-linecap="round" stroke-linejoin="round" d="m 223.95413,15.996971 a 15.992355,15.992355 0 0 1 15.99236,15.992354"/><path id="galaDisplay6" stroke-linecap="butt" stroke-linejoin="miter" stroke-opacity="1" d="M 223.95413,191.91287 H 32.04587"/><path id="galaDisplay7" stroke-linecap="round" stroke-linejoin="round" d="m 223.95413,191.91287 a 15.992354,15.992354 0 0 0 15.99236,-15.99235"/><path id="galaDisplay8" stroke-linecap="round" stroke-linejoin="round" d="M 32.04587,191.91287 A 15.992354,15.992354 0 0 1 16.053504,175.92052"/><path id="galaDisplay9" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 191.96942,239.88994 H 64.030574"/><path id="galaDisplaya" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 143.99235,191.91288 v 47.97706"/><path id="galaDisplayb" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="m 112.00764,191.91288 v 47.97706"/><path id="galaDisplayc" stroke-linecap="round" stroke-linejoin="round" stroke-opacity="1" d="M 16.053512,159.92816 H 239.94649"/></g>`),
		g.Group(children),
	)
}