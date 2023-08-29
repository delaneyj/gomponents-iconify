package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreCircleTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.5 14a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Zm6 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0ZM14 2C7.373 2 2 7.373 2 14s5.373 12 12 12s12-5.373 12-12S20.627 2 14 2ZM3.5 14C3.5 8.201 8.201 3.5 14 3.5S24.5 8.201 24.5 14S19.799 24.5 14 24.5S3.5 19.799 3.5 14Z"/>`),
		g.Group(children),
	)
}