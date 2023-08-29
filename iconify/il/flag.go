package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M628 340q10 11 5 25t-22 14H182q-11 0-18 8t-5 19l56 316q2 11-5 19t-18 9h-45q-20 0-23-20L1 35q-2-11 5-19t18-8h535q14 0 21 11t-1 24l-78 136q-9 15 3 27z"/>`),
		g.Group(children),
	)
}