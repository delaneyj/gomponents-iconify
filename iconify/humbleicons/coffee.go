package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12H4v4a4 4 0 0 0 4 4h5a4 4 0 0 0 4-4v-4zm0 0h2a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-2m-4-8s1-1 .5-2l-1-2C12 4 13 3 13 3M8.64 9s1-1 .5-2l-1-2c-.5-1 .5-2 .5-2"/>`),
		g.Group(children),
	)
}