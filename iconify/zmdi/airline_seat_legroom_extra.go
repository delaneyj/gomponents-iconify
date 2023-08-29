package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomExtra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 192q0 27 18.5 45.5T107 256h128v43H107q-44 0-75.5-31.5T0 192V0h43v192zm401 112q7 12 2.5 25T429 348l-79 36l-73-149H128q-27 0-45.5-19T64 171V0h128v128h75q26 0 38 24l73 149l23-11q12-5 24.5-1.5T444 304z"/>`),
		g.Group(children),
	)
}