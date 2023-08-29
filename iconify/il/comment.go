package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M718 5q10 0 17 7t6 17v657q0 5-3 6t-7-1L579 538H23q-10 0-16-7t-7-16V29q0-10 7-17t16-7h695z"/>`),
		g.Group(children),
	)
}