package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DividerTallTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 1a.5.5 0 0 1 .5.5v17a.5.5 0 0 1-1 0v-17a.5.5 0 0 1 .5-.5Z"/>`),
		g.Group(children),
	)
}