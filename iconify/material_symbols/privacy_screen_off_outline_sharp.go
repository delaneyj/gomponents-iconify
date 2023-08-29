package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivacyScreenOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.45 23.3l-3.3-3.3H2V5.075h.225L.65 3.5l1.425-1.425l19.8 19.8ZM4 16.6l4.875-4.875l-2.075-2.1l-2.8 2.8Zm0-7l1.375 1.375L4 9.6ZM5.4 18h9.75l-4.875-4.875ZM22 19.175l-2-2.025V6h-2.6l-4.275 4.275l-1.4-1.4L14.575 6H10.4l-.75.8L6.85 4H22Zm-12.425-6.75Zm4.85-.85Z"/>`),
		g.Group(children),
	)
}