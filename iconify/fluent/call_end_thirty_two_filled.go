package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CallEndThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.557 10.724C9.054 9.621 12.385 9 15.997 9c3.618 0 6.968.623 9.474 1.73c2.439 1.078 4.402 2.776 4.522 5.084a2.78 2.78 0 0 1-.052.66l-.502 2.67a3.5 3.5 0 0 1-4.063 2.797l-3.23-.584a3 3 0 0 1-2.403-2.342l-.614-2.955a10.089 10.089 0 0 0-6.266 0l-.614 2.955a3 3 0 0 1-2.403 2.342l-3.23.584a3.5 3.5 0 0 1-4.062-2.797l-.499-2.65a2.791 2.791 0 0 1-.049-.722c.168-2.285 2.116-3.972 4.551-5.048Z"/>`),
		g.Group(children),
	)
}