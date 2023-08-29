package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassHalf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 4v2h2v4a7.001 7.001 0 0 0 3.406 6A7.001 7.001 0 0 0 9 22v4H7v2h18v-2h-2v-4a7.001 7.001 0 0 0-3.406-6A7.001 7.001 0 0 0 23 10V6h2V4zm4 2h10v4c0 2.773-2.227 5-5 5s-5-2.227-5-5zm1.156 5c.446 1.723 1.98 3 3.844 3c1.863 0 3.398-1.277 3.844-3zM16 17c2.773 0 5 2.227 5 5v4h-1c0-2.21-1.79-4-4-4s-4 1.79-4 4h-1v-4c0-2.773 2.227-5 5-5z"/>`),
		g.Group(children),
	)
}