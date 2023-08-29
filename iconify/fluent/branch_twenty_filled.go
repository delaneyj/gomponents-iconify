package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 5a3 3 0 0 1-2.145 2.876c.178.223.416.483.72.744c.726.622 1.834 1.252 3.425 1.363a3 3 0 1 1 .168 1.01c-1.96-.078-3.342-.841-4.243-1.613a6.62 6.62 0 0 1-.425-.398v3.06a3 3 0 1 1-1 0V7.958A3 3 0 1 1 10 5Z"/>`),
		g.Group(children),
	)
}