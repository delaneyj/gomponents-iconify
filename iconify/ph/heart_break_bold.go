package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBreakBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M178 28a65.6 65.6 0 0 0-46.67 19.33L128 50.66l-3.33-3.33A66 66 0 0 0 12 94c0 72.34 105.81 130.14 110.31 132.57a12 12 0 0 0 11.38 0C138.19 224.14 244 166.34 244 94a66.08 66.08 0 0 0-66-66Zm-5.49 142.36a328.69 328.69 0 0 1-44.51 31.8a328.69 328.69 0 0 1-44.51-31.8C61.82 151.77 36 123.42 36 94a42 42 0 0 1 71.7-29.7l3.3 3.33l-7.52 7.52a12 12 0 0 0 0 17l21.7 21.7l-13.7 13.69a12 12 0 0 0 17 17l22.18-22.19a12 12 0 0 0 0-17L129 83.63l19.3-19.33A42 42 0 0 1 220 94c0 29.42-25.82 57.77-47.49 76.36Z"/>`),
		g.Group(children),
	)
}