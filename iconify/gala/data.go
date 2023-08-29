package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Data(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaData0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-width="16"><path id="galaData1" d="M 239.98507,55.993592 A 111.98507,39.994664 0 0 1 128,95.988256 111.98507,39.994664 0 0 1 16.01493,55.993592 111.98507,39.994664 0 0 1 128,15.998927 111.98507,39.994664 0 0 1 239.98507,55.993592 Z"/><path id="galaData2" d="m 239.98507,199.97441 a 111.98507,39.994664 0 0 1 -55.99253,34.63639 111.98507,39.994664 0 0 1 -111.985079,0 111.98507,39.994664 0 0 1 -55.992531,-34.6364"/><path id="galaData3" d="m 239.98507,151.9808 a 111.98507,39.994664 0 0 1 -55.99253,34.6364 111.98507,39.994664 0 0 1 -111.985079,-1e-5 A 111.98507,39.994664 0 0 1 16.01493,151.9808"/><path id="galaData4" d="m 239.98507,103.9872 a 111.98507,39.994664 0 0 1 -55.99253,34.6364 111.98507,39.994664 0 0 1 -111.985079,0 111.98507,39.994664 0 0 1 -55.992531,-34.6364"/><path id="galaData5" stroke-opacity="1" d="M 16.01493,55.99377 V 199.97441"/><path id="galaData6" stroke-opacity="1" d="M 239.98507,55.993592 V 199.97441"/></g>`),
		g.Group(children),
	)
}