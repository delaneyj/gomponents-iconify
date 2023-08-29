package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func XingLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.444 3.5L13.81 14.99L17.857 22h-2.31l-4.045-7.009H11.5L18.134 3.5h2.31ZM8.31 7l2.422 4.196l-.002.001L7.67 16.5H5.36l3.061-5.305L6.002 7H8.31Z"/>`),
		g.Group(children),
	)
}