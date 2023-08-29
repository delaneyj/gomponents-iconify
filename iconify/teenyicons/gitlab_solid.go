package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitlabSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M2.974.342a.5.5 0 0 0-.96.037l-2 8a.5.5 0 0 0 .16.5l7 6a.5.5 0 0 0 .651 0l7-6a.5.5 0 0 0 .16-.5l-2-8a.5.5 0 0 0-.96-.037L10.14 6H4.86L2.974.342Z"/>`),
		g.Group(children),
	)
}