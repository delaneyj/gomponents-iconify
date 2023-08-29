package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M3 1L0 3.5L3 6V4h2v2l3-2.5L5 1v2H3V1z"/>`),
		g.Group(children),
	)
}