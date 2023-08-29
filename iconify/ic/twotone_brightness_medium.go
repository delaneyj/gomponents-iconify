package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneBrightnessMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-opacity=".3" d="M18 9.52V6h-3.52L12 3.52L9.52 6H6v3.52L3.52 12L6 14.48V18h3.52L12 20.48L14.48 18H18v-3.52L20.48 12L18 9.52zM12 18V6c3.31 0 6 2.69 6 6s-2.69 6-6 6z"/><path fill="currentColor" d="M20 8.69V4h-4.69L12 .69L8.69 4H4v4.69L.69 12L4 15.31V20h4.69L12 23.31L15.31 20H20v-4.69L23.31 12L20 8.69zm-2 5.79V18h-3.52L12 20.48L9.52 18H6v-3.52L3.52 12L6 9.52V6h3.52L12 3.52L14.48 6H18v3.52L20.48 12L18 14.48zM12 6v12c3.31 0 6-2.69 6-6s-2.69-6-6-6z"/>`),
		g.Group(children),
	)
}