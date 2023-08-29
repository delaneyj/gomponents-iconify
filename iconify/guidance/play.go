package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m2 22.5l1 .5c7.5-5 12.5-7.75 19-10.25v-1.5C15.5 8.75 10.5 6 3 1l-1 .5s.5 6.2.5 10.5S2 22.5 2 22.5Z"/>`),
		g.Group(children),
	)
}