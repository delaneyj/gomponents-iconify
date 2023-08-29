package gala

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Editor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g id="galaEditor0" fill="none" stroke="currentColor" stroke-dasharray="none" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="4" stroke-opacity="1" stroke-width="16"><path id="galaEditor1" d="m 16,64 224.93778,0.09256"/><path id="galaEditor2" d="M 48,16 H 207.91114 C 225.62929,16 240,30.281849 240,48 v 160 c 0,17.71816 -14.28185,32 -32,32 H 48 C 30.281848,240 16.069099,225.73073 16.06221,208.01257 L 16,48 C 15.993112,30.281851 30.281848,16 48,16 Z"/><path id="galaEditor3" d="M 191.96444,64.092555 192,16"/><path id="galaEditor4" d="M 48.044437,112.06589 H 80.02666"/><path id="galaEditor5" d="M 48.044437,144.04812 H 175.97333"/><path id="galaEditor6" d="M 48.044437,176.03034 H 127.99999"/><path id="galaEditor7" d="M 48.044437,208.01256 H 80.02666"/></g>`),
		g.Group(children),
	)
}