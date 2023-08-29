package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicroservicesOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11 21l-4-2.2v-4.5l4-2.2l4 2.2v4.5L11 21zm-2-3.4l2 1.1l2-1.1v-2.2l-2-1.1l-2 1.1v2.2zM6 30l-4-2.2v-4.5L6 21l4 2.2v4.5L6 30zm-2-3.4l2 1.1l2-1.1v-2.2l-2-1.1l-2 1.1v2.2zM16 30l-4-2.2v-4.5l4-2.2l4 2.2v4.5L16 30zm-2-3.4l2 1.1l2-1.1v-2.2l-2-1.1l-2 1.1v2.2zM26 30l-4-2.2v-4.5l4-2.2l4 2.2v4.5L26 30zm-2-3.4l2 1.1l2-1.1v-2.2l-2-1.1l-2 1.1v2.2zm.6-15L22 14.2V9.4l3-1.7V3.2L21 1l-4 2.2v4.5l3 1.7v4.7l-2.6-2.6L16 13l5 5l5-5l-1.4-1.4zM19 4.4l2-1.1l2 1.1v2.2l-2 1.1l-2-1.1V4.4z"/>`),
		g.Group(children),
	)
}