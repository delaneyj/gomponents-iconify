package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M13 3L7 1L1 3v16l6-2l6 2l6-2V1zM7 14.89l-4 1.36V4.35L7 3zm10 .75L13 17V5.1l4-1.36z"/>`),
		g.Group(children),
	)
}