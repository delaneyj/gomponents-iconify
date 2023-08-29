package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberZeroSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-14Zm7 6.124C10 8.913 10.934 8 12 8c1.066 0 2 .913 2 2.125v3.75C14 15.088 13.066 16 12 16c-1.066 0-2-.912-2-2.125v-3.75ZM12 6c-2.247 0-4 1.886-4 4.125v3.75C8 16.115 9.753 18 12 18s4-1.886 4-4.125v-3.75C16 7.885 14.247 6 12 6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}