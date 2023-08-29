package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatLegroomReduced(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M362 346q3 15-6.5 26.5T331 384h-96v-64l21-85H128q-26 0-45-19t-19-45V0h128v128h107q17 0 29.5 12.5T341 171l-42 149h30q12 0 21.5 7t11.5 19zM43 192q0 27 18.5 45.5T107 256h85v43h-85q-44 0-75.5-31.5T0 192V0h43v192z"/>`),
		g.Group(children),
	)
}