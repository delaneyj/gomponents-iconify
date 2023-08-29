package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17v-2a9 9 0 1 1 18 0v2M6 17v-2a6 6 0 0 1 12 0v2m-9 0v-2a3 3 0 1 1 6 0v2"/>`),
		g.Group(children),
	)
}