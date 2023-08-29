package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func House(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M733 289q8 6 8 18v406q0 10-6 16t-17 7H487q-10 0-17-7t-7-16V458q0-10-6-17t-17-6H301q-10 0-16 6t-7 17v255q0 10-6 16t-17 7H23q-10 0-16-7t-7-16V307q0-12 9-18L356 7q15-12 29 0z"/>`),
		g.Group(children),
	)
}