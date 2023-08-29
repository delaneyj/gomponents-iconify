package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m19 10l-4-3v2h-4V5h2l-3-4l-3 4h2v4H5V7l-4 3l4 3v-2h4v4H7l3 4l3-4h-2v-4h4v2z"/>`),
		g.Group(children),
	)
}