package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrendingUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 7l-5.846 5.938c-.105.106-.158.16-.205.202c-.76.68-1.907.68-2.667 0a4.814 4.814 0 0 1-.205-.203c-.105-.106-.158-.16-.205-.202a2 2 0 0 0-2.667 0a4.86 4.86 0 0 0-.204.202L4 17M20 7v6m0-6h-6"/>`),
		g.Group(children),
	)
}