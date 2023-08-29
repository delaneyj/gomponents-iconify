package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendCancelled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 19l4 4m0 0l4 4m-4-4l-4 4m4-4l4-4m1-16L3 15l12 2.5M29 3L15 17.5M29 3l-4.375 11.375M15 17.5l.625 2.875"/>`),
		g.Group(children),
	)
}