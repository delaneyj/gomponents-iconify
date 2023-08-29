package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchForkSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 3.5a2.5 2.5 0 0 1-2 2.45V8h4.5A1.5 1.5 0 0 0 11 6.5v-.55a2.5 2.5 0 1 1 1 0v.55A2.5 2.5 0 0 1 9.5 9H5v1.05a2.5 2.5 0 1 1-1 0v-4.1A2.5 2.5 0 1 1 7 3.5Z"/>`),
		g.Group(children),
	)
}