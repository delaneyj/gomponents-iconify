package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mathematics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2H4l5 8l-5 8h12v-4h-2v2H8.25L12 10L8.25 4H14v2h2V2z"/>`),
		g.Group(children),
	)
}