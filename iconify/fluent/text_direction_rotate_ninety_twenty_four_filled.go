package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirectionRotateNinetyTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M20.385 8.673l-9 3.75a1 1 0 1 1-.77-1.846L13 9.584V5.917l-2.385-.994a1 1 0 0 1 .77-1.846l9 3.75a1 1 0 0 1 0 1.846zM15 8.75l2.4-1l-2.4-1v2z" fill="currentColor"/><path d="M8 4a1 1 0 0 0-2 0v13.586l-.293-.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l2-2a1 1 0 0 0-1.414-1.414L8 17.586V4z" fill="currentColor"/><path d="M16 13a1 1 0 0 0-1 1v3.586l-.293-.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0l2-2a1 1 0 0 0-1.414-1.414l-.293.293V14a1 1 0 0 0-1-1z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}