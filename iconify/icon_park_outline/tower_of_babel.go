package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TowerOfBabel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="M20 14.5V4h8v9.5M14 25v-9.538L34 13v10M11 35v-9l26-3v9m3 12H8v-8l32-4v12Z"/>`),
		g.Group(children),
	)
}