package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoFixOffSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20 7l-.95-2.05L17 4l2.05-.95L20 1l.95 2.05L23 4l-2.05.95L20 7Zm-4.45 5.7L11.3 8.45l2.875-2.875l4.25 4.25L15.55 12.7Zm4.25 9.9l-7.1-7.05l-6.875 6.875L1.6 18.15l6.85-6.85L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4Z"/>`),
		g.Group(children),
	)
}