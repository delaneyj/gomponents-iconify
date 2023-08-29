package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UsbFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l3 5h-2v7.381l3-1.499l-.001-.882H15V7h4v4h-1.001L18 13.118l-5 2.5v1.553A3.001 3.001 0 0 1 12 23a3 3 0 0 1-1.31-5.7L6 14l-.001-2.268a2 2 0 1 1 2.001 0V13l3 2.086V6H9l3-5Z"/>`),
		g.Group(children),
	)
}