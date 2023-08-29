package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignBottomTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.75 25a.75.75 0 0 1 0-1.5h22.5a.75.75 0 0 1 0 1.5H2.75ZM5 19.25A2.75 2.75 0 0 0 7.75 22h2.5A2.75 2.75 0 0 0 13 19.25V5.75A2.75 2.75 0 0 0 10.25 3h-2.5A2.75 2.75 0 0 0 5 5.75v13.5Zm10 0A2.75 2.75 0 0 0 17.75 22h2.5A2.75 2.75 0 0 0 23 19.25v-8a2.75 2.75 0 0 0-2.75-2.75h-2.5A2.75 2.75 0 0 0 15 11.25v8Z"/>`),
		g.Group(children),
	)
}