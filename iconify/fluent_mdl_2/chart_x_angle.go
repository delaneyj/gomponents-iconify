package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartXAngle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 347v1571q75-6 151-25t148-51t135-76t115-102H640v-128h512v512h-128v-293l-19 18q-143 137-315 206t-370 69h-64V37l1645 1646l-90 90L384 347z"/>`),
		g.Group(children),
	)
}