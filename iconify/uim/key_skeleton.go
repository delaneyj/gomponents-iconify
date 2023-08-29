package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySkeleton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 12a5 5 0 1 0 0 10a5 5 0 0 0 0-10zm-1.415 7.413a1 1 0 1 1 0-2a1 1 0 0 1 0 2z" opacity=".5"/><path fill="currentColor" d="m21.001 4.413l.706-.706a1 1 0 1 0-1.414-1.414L9.754 12.832a5.02 5.02 0 0 1 1.414 1.414l5.591-5.591l2.12 2.12a1 1 0 0 0 1.414-1.414l-2.12-2.12l1.414-1.414l.706.705a1 1 0 0 0 1.414-1.414l-.706-.705z"/>`),
		g.Group(children),
	)
}