package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChefHatOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.72 4.712a4 4 0 0 1 7.19 1.439A4 4 0 0 1 18 13.874V14m0 4v3H6v-7.126a4 4 0 0 1 .081-7.796m.08 10.931L17 17M3 3l18 18"/>`),
		g.Group(children),
	)
}