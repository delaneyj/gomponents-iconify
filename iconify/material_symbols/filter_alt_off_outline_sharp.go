package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterAltOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.8 11.975l-1.425-1.425L16.95 6H8.825l-2-2H21.05l-6.25 7.975ZM19.775 22.6L14 16.825V20h-4v-7.175l-8.6-8.6L2.8 2.8l18.4 18.4l-1.425 1.4Zm-6.4-12.05Z"/>`),
		g.Group(children),
	)
}