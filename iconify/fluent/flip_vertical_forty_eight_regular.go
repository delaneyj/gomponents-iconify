package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalFortyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M39.423 4.197c.36.23.577.627.577 1.053v15.5c0 .69-.56 1.25-1.25 1.25H5.25a1.25 1.25 0 0 1-.525-2.384l33.5-15.5a1.25 1.25 0 0 1 1.198.08ZM10.928 19.5H37.5V7.206L10.928 19.5ZM40 43a1 1 0 0 1-1.425.905l-34-16A1 1 0 0 1 5 26h34a1 1 0 0 1 1 1v16Z"/>`),
		g.Group(children),
	)
}