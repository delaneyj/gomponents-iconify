package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchToPageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.11 0-2 .89-2 2v16c0 1.11.89 2 2 2h16c1.11 0 2-.89 2-2V4c0-1.11-.89-2-2-2m0 18H4V4h16M8.83 10.25L7.29 8.71L6 10V6h4L8.71 7.29l1.54 1.54m3.5 0l1.54-1.54L14 6h4v4l-1.29-1.29l-1.54 1.54m0 3.5l1.54 1.54L18 14v4h-4l1.29-1.29l-1.54-1.54m-3.5 0l-1.54 1.54L10 18H6v-4l1.29 1.29l1.54-1.54"/>`),
		g.Group(children),
	)
}