package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etsy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M518 183v655q103 1 191.5-1.5T835 831l37-3q68-2 90.5-24.5T1002 709l33-142h103l-14 322l7 319h-103l-29-127q-15-68-45-93t-84-26q-87-8-352-8v556q0 78 43.5 115.5T695 1663h357q35 0 59.5-2t55-7.5t54-18t48.5-32t46-50.5t39-73l93-216h89q-6 37-31.5 252t-30.5 276q-146-5-263.5-8t-162.5-4H376L0 1792v-102l127-25q67-13 91.5-37t25.5-79l8-643q3-402-8-645q-2-61-25.5-84T127 141L0 117V15l376 12h702q139 0 374-27q-6 68-14 194.5T1426 414l-5 92h-93l-32-124q-31-121-74-179.5T1109 144H561q-28 0-35.5 8.5T518 183z"/>`),
		g.Group(children),
	)
}