package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaMouse0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaMouse1" stroke-opacity="1" d="m 63.998705,64.001121 c 0,64.001289 -16.00032,48.000969 -16.00032,96.001929"/><path id="galaMouse2" stroke-opacity="1" d="m 192.00129,64.001121 c 0,-16.000322 0,-22.762857 -16.00032,-32.000645"/><path id="galaMouse3" stroke-opacity="1" d="m 128,48.000796 v 32.00065"/><path id="galaMouse4" stroke-opacity="1" d="m 79.99903,32.000476 c 0,0 32.00064,-16.000321 48.00097,-16.000321 16.00032,0 48.00097,16.000321 48.00097,16.000321"/><path id="galaMouse5" stroke-opacity="1" d="m 192.00129,64.001121 c 0,64.001289 16.00032,48.000969 16.00032,96.001929"/><path id="galaMouse6" d="m 208.00161,160.00305 a 80.00161,80.00161 0 0 1 -40.00081,69.28343 80.00161,80.00161 0 0 1 -80.001607,0 80.00161,80.00161 0 0 1 -40.000803,-69.28343"/><path id="galaMouse7" stroke-opacity="1" d="m 63.998705,64.001125 c 0,-16.000326 2e-6,-22.762857 16.000325,-32.000649"/></g>`),
		g.Group(children),
	)
}