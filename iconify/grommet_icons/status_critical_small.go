package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusCriticalSmall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" stroke="currentColor" d="M6.712 1.263a1.005 1.005 0 0 0-1.424 0L1.263 5.288a1.005 1.005 0 0 0 0 1.424l4.025 4.025a1.005 1.005 0 0 0 1.424 0l4.025-4.025a1.005 1.005 0 0 0 0-1.424L6.712 1.263Z"/>`),
		g.Group(children),
	)
}