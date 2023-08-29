package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 15a7 7 0 0 0 4.2-1.4l5.4 5.4l1.4-1.4l-5.4-5.4A7 7 0 1 0 8 15Zm0-2A5 5 0 1 1 8 3a5 5 0 0 1 0 10Zm1-6h2v2H9v2H7V9H5V7h2V5h2Z"/>`),
		g.Group(children),
	)
}