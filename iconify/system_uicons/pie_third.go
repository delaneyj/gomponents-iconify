package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieThird(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.519 2.747a8.013 8.013 0 0 0-5.791 5.849M10.5 2.5a8 8 0 1 1 0 16c-4.418 0-8-3.5-8-8h8v-8z"/>`),
		g.Group(children),
	)
}