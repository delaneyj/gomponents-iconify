package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Location(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M671 220q7 3 12 8t5 13t-3 13t-11 9L396 389q-8 5-11 12L259 678q-3 8-9 11t-14 3t-12-4t-9-12L2 36q-5-14 6-24t24-5z"/>`),
		g.Group(children),
	)
}