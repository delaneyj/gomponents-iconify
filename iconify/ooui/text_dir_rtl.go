package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDirRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m1 10l6-5v4h7v2H7v4zm13-8V1h1.5a1.49 1.49 0 0 1 1 .39a1.49 1.49 0 0 1 1-.39H19v1h-1.5a.5.5 0 0 0-.5.5v15a.5.5 0 0 0 .5.5H19v1h-1.5a1.49 1.49 0 0 1-1-.39a1.49 1.49 0 0 1-1 .39H14v-1h1.5a.5.5 0 0 0 .5-.5v-15a.5.5 0 0 0-.5-.5z"/>`),
		g.Group(children),
	)
}