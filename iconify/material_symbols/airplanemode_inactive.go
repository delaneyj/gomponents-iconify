package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplanemodeInactive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 21l-3.5 1v-1.5l2-1.5v-5.5L2 16v-2l5.8-3.4l-6.4-6.4l1.4-1.4l18.4 18.4l-1.4 1.4l-6.3-6.3V19l2 1.5V22L12 21Zm0-19q.625 0 1.063.438T13.5 3.5V9l8.5 5v2l-4.45-1.325L10.5 7.65V3.5q0-.625.438-1.063T12 2Z"/>`),
		g.Group(children),
	)
}