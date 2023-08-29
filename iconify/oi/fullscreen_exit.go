package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1 0L0 1l1.5 1.5L0 4h4V0L2.5 1.5L1 0zm3 4v4l1.5-1.5L7 8l1-1l-1.5-1.5L8 4H4z"/>`),
		g.Group(children),
	)
}