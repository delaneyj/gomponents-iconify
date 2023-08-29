package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10v9a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-9M6 10l6-6l6 6M6 10l-2 2m14-2l2 2m-10 1h4v4h-4v-4z"/>`),
		g.Group(children),
	)
}