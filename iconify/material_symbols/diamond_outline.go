package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21L2 9l3-6h14l3 6l-10 12ZM9.625 8h4.75l-1.5-3h-1.75l-1.5 3ZM11 16.675V10H5.45L11 16.675Zm2 0L18.55 10H13v6.675ZM16.6 8h2.65l-1.5-3H15.1l1.5 3ZM4.75 8H7.4l1.5-3H6.25l-1.5 3Z"/>`),
		g.Group(children),
	)
}