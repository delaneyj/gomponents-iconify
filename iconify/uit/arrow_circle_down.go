package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.146 11.646L12.5 15.293V7.5a.5.5 0 0 0-1 0v7.793l-3.647-3.646a.5.5 0 0 0-.707.707l4.5 4.5A.5.5 0 0 0 12 17c.011 0 .02-.005.03-.006a.498.498 0 0 0 .163-.033a.5.5 0 0 0 .162-.109l4.498-4.498a.5.5 0 0 0-.707-.708zM12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10zm0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9z"/>`),
		g.Group(children),
	)
}