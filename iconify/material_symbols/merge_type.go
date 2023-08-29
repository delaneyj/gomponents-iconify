package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.6 20L11 14.4V6.875l-2.6 2.6L6.975 8.05L12 3.025l5 5l-1.425 1.425L13 6.875V13.6l5 5l-1.4 1.4Zm-9.2.025l-1.4-1.4l3.175-3.2L10.6 16.85l-3.2 3.175Z"/>`),
		g.Group(children),
	)
}