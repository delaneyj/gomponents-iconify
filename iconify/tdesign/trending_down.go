package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.086 7.5L8.5 13.914l4-4L17.586 15H13.5v2H21V9.5h-2v4.086l-6.5-6.5l-4 4l-5-5L2.086 7.5Z"/>`),
		g.Group(children),
	)
}