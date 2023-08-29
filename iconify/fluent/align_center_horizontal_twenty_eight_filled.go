package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterHorizontalTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M26 13.75a.75.75 0 0 1-.75.75H23V18a2.75 2.75 0 0 1-2.75 2.75h-2.5A2.75 2.75 0 0 1 15 18v-3.5h-2v6.25a2.75 2.75 0 0 1-2.75 2.75h-2.5A2.75 2.75 0 0 1 5 20.75V14.5H2.75a.75.75 0 0 1 0-1.5H5V7.25A2.75 2.75 0 0 1 7.75 4.5h2.5A2.75 2.75 0 0 1 13 7.25V13h2v-3a2.75 2.75 0 0 1 2.75-2.75h2.5A2.75 2.75 0 0 1 23 10v3h2.25a.75.75 0 0 1 .75.75Z"/>`),
		g.Group(children),
	)
}