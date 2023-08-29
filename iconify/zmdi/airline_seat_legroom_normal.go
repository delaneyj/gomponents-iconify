package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomNormal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M43 192q0 27 18.5 45.5T107 256h128v43H107q-44 0-75.5-31.5T0 192V0h43v192zm330 128q14 0 23 9.5t9 22.5t-9 22.5t-23 9.5h-96V235H128q-26 0-45-19t-19-45V0h128v128h107q17 0 29.5 12.5T341 171v149h32z"/>`),
		g.Group(children),
	)
}