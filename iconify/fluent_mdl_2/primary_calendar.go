package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrimaryCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2031 1537l-266 197l102 313l-267-193l-266 194l101-314l-266-197h335l97-310l95 310h335zm17-1409v1249h-128V640H128v1152h1116l4 3l-41 125H0V128h384V0h128v128h1024V0h128v128h384zm-128 384V256h-256v128h-128V256H512v128H384V256H128v256h1792z"/>`),
		g.Group(children),
	)
}