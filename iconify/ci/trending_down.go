package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 17l-5.846-5.938a4.945 4.945 0 0 0-.205-.202a1.999 1.999 0 0 0-2.667 0c-.047.042-.1.096-.205.203c-.105.106-.158.16-.205.202a2 2 0 0 1-2.667 0A4.898 4.898 0 0 1 8 11.062L4 7m16 10v-6m0 6h-6"/>`),
		g.Group(children),
	)
}