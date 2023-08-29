package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseWithShieldOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q1.725-.425 2.863-1.988T16 12.55V10l-4-2l-4 2v2.55q0 1.9 1.137 3.463T12 18Zm-8 3V9l8-6l8 6v12H4Zm2-2h12v-9l-6-4.5L6 10v9Zm6-6.75Z"/>`),
		g.Group(children),
	)
}