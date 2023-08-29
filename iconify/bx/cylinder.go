package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cylinder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22c5.131 0 9-1.935 9-4.5V7h-.053c.033-.164.053-.33.053-.5C21 3.935 17.131 2 12 2C7.209 2 3.52 3.688 3.053 6H3v11.5c0 2.565 3.869 4.5 9 4.5zm0-2c-4.273 0-7-1.48-7-2.5V9.394C6.623 10.387 9.111 11 12 11s5.377-.613 7-1.606V17.5c0 1.02-2.727 2.5-7 2.5zm0-16c4.273 0 7 1.48 7 2.5S16.273 9 12 9S5 7.52 5 6.5S7.727 4 12 4z"/>`),
		g.Group(children),
	)
}