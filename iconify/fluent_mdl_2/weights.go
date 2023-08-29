package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Weights(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1883 640l128 128l-283 282l-320-319l-677 677l319 320l-282 283l-128-128l-160 159l-192-191q-10 9-20 21t-22 22t-25 18t-29 8q-26 0-45-19t-19-45q0-15 7-28t18-25t23-22t21-21L6 1568l159-160l-128-128l283-282l320 319l677-677l-319-320l282-283l128 128L1568 6l192 191q10-9 20-21t22-22t25-18t29-8q26 0 45 19t19 45q0 15-7 28t-18 25t-23 22t-21 21l191 192l-159 160zM549 1792l-293-293l-69 69l293 293l69-69zm321-64l-550-550l-101 102l549 549l102-101zm629-1472l293 293l69-69l-293-293l-69 69zm330 512l-549-549l-102 101l550 550l101-102z"/>`),
		g.Group(children),
	)
}