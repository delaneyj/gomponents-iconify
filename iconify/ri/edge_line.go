package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.008 14.001A4.566 4.566 0 0 0 8 14.25C8 16.632 9.753 19 13 19c2.373 0 4.528-.655 6-1.553v3.35C17.211 21.564 15.112 22 13 22c-5.502 0-8-3.47-8-7.75c0-3.231 2.041-6 4.943-7.164C8.54 8.663 8 10.341 8 10.996L18 11c0-3.406-2.548-6-6-6c-5 0-8.001 3.988-9 5.999C3.29 6.237 7.01 2 12 2c5.2 0 9 4.03 9 9v3H8l.007.001Z"/>`),
		g.Group(children),
	)
}