package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.086 16.5L8.5 10.086l4 4L17.586 9H13.5V7H21v7.5h-2v-4.086l-6.5 6.5l-4-4l-5 5L2.086 16.5Z"/>`),
		g.Group(children),
	)
}