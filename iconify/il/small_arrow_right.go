package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M23 44q10 0 16 6l163 162q7 8 7 17t-7 16L39 407q-7 7-16 7t-16-7t-7-16t7-16l146-146L7 83q-7-7-7-17t6.5-16T23 44z"/>`),
		g.Group(children),
	)
}