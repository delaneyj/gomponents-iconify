package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrikethroughA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6 8h2.6L10 4l1.4 4h2.7l-2.4-6H8.3zm-5 2v2h3.4L2 18h3l1.7-4.6h6.6L15 18h3l-2.4-6H19v-2z"/>`),
		g.Group(children),
	)
}