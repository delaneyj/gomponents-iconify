package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseWithShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q1.725-.425 2.863-1.988T16 12.55V10l-4-2l-4 2v2.55q0 1.9 1.137 3.463T12 18Zm-8 3V9l8-6l8 6v12H4Z"/>`),
		g.Group(children),
	)
}