package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActionRedo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3.5 1C1.57 1 0 2.57 0 4.5a2.5 2.5 0 0 1 5 0V5H4l2 2l2-2H7v-.5C7 2.57 5.43 1 3.5 1z"/>`),
		g.Group(children),
	)
}