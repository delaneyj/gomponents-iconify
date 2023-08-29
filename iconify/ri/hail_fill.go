package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HailFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.995 17.794a4 4 0 0 0-5.085-3.644A4.001 4.001 0 0 0 6 15c0 1.08.428 2.059 1.122 2.778a8 8 0 1 1 9.335-10.68a5.5 5.5 0 0 1 2.537 10.696ZM10 17a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm5 3a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm-5 3a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}