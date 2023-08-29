package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximizeFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 11.5A5.5 5.5 0 0 1 11.5 6h25a5.5 5.5 0 0 1 5.5 5.5v25a5.5 5.5 0 0 1-5.5 5.5h-25A5.5 5.5 0 0 1 6 36.5v-25ZM11.5 9A2.5 2.5 0 0 0 9 11.5v25a2.5 2.5 0 0 0 2.5 2.5h25a2.5 2.5 0 0 0 2.5-2.5v-25A2.5 2.5 0 0 0 36.5 9h-25Z"/>`),
		g.Group(children),
	)
}