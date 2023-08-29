package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TvTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.75 4A2.75 2.75 0 0 0 2 6.75v8.5A2.75 2.75 0 0 0 4.75 18h14.5A2.75 2.75 0 0 0 22 15.25v-8.5A2.75 2.75 0 0 0 19.25 4H4.75ZM5 20.25a.75.75 0 0 1 .75-.75h12.5a.75.75 0 0 1 0 1.5H5.75a.75.75 0 0 1-.75-.75Z"/>`),
		g.Group(children),
	)
}