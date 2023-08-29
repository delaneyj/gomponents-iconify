package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 11l-2 4l3.408 1.363a4 4 0 0 1 2.229 2.229L9 22l4-2l-1.21-2.42a2 2 0 0 1 .679-2.56L14 14l4 7l3-4l-2.29-7.469l.715-.714c1.412-1.412 2.71-3.682 1.075-5.317c-1.635-1.635-3.91-.34-5.316 1.077l-.72.708L7 3L3 6l7 4l-1.02 1.531a2 2 0 0 1-2.56.68L4 11Z"/>`),
		g.Group(children),
	)
}