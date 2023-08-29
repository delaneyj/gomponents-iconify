package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnfoldLessHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.59 5.41L15.17 4L12 7.17L8.83 4L7.41 5.41L12 10m-4.59 8.59L8.83 20L12 16.83L15.17 20l1.41-1.41L12 14l-4.59 4.59Z"/>`),
		g.Group(children),
	)
}