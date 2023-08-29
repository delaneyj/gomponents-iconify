package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeBoth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="m4 0l1.66 1.66l-4 4L0 4v4h4L2.34 6.34l4-4L8 4V0H4z"/>`),
		g.Group(children),
	)
}