package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaBag0" fill="none" stroke-width="16"><path id="galaBag1" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="m 64,80 h 128 c 16,0 29.33333,16 32,32 l 16,96 c 2.66807,16.00842 -16,32 -32,32 H 48 C 32,240 13.33193,224.00842 16,208 L 32,112 C 34.666667,96 48,80 64,80 Z"/><path id="galaBag2" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="M 80,112 V 63.814079"/><path id="galaBag3" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="m 176,64 v 48"/><path id="galaBag4" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" d="M 19.090159,191.31828 H 236.90984"/><path id="galaBag5" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" d="M 176,64 C 176,48 166.70076,30.94703 151.90703,22.405869 137.1133,13.86471 118.88668,13.86471 104.09296,22.40587 89.299233,30.947031 80.000002,48 80,64"/><rect id="galaBag6" width="80" height="16" x="16" y="240" stroke="none"/></g>`),
		g.Group(children),
	)
}