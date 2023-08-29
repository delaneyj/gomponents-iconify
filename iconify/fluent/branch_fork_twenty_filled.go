package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchForkTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 5a3 3 0 0 1-2.5 2.959V10H12a1.5 1.5 0 0 0 1.5-1.5v-.541a3 3 0 1 1 1 0V8.5A2.5 2.5 0 0 1 12 11H6.5v1.041a3 3 0 1 1-1 0V7.96A3 3 0 1 1 9 5Z"/>`),
		g.Group(children),
	)
}