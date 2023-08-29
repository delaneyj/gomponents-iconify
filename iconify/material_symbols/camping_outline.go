package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CampingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-4.65l8.75-11.8L9 3.2L10.6 2L12 3.875L13.4 2L15 3.2l-1.75 2.35L22 17.35V22H2ZM12 7.225L4 18v2h3l5-7l5 7h3v-2L12 7.225ZM9.45 20h5.1L12 16.45L9.45 20ZM12 13l5 7l-5-7l-5 7l5-7Z"/>`),
		g.Group(children),
	)
}