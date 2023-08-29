package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MatchAppLayoutTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 8c.966 0 1.75.784 1.75 1.75v4.5A1.75 1.75 0 0 1 9.5 16H3.75A1.75 1.75 0 0 1 2 14.25v-4.5a1.75 1.75 0 0 1 1.606-1.744L3.75 8H9.5Zm10.75 0c.966 0 1.75.784 1.75 1.75v4.5A1.75 1.75 0 0 1 20.25 16H14.5a1.75 1.75 0 0 1-1.75-1.75v-4.5c0-.966.784-1.75 1.75-1.75h5.75Z"/>`),
		g.Group(children),
	)
}