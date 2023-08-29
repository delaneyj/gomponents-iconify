package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeftAndRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M42 6V42M16 20L12 24L16 28M32 20L36 24L32 28M24 6L24 42M6 6L6 42"/>`),
		g.Group(children),
	)
}