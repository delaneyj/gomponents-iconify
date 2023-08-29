package zondicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.8 15.8L0 13v7h7l-2.8-2.8l4.34-4.32l-1.42-1.42L2.8 15.8zM17.2 4.2L20 7V0h-7l2.8 2.8l-4.34 4.32l1.42 1.42L17.2 4.2zm-1.4 13L13 20h7v-7l-2.8 2.8l-4.32-4.34l-1.42 1.42l4.33 4.33zM4.2 2.8L7 0H0v7l2.8-2.8l4.32 4.34l1.42-1.42L4.2 2.8z"/>`),
		g.Group(children),
	)
}