package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Screenshot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 19a2 2 0 0 1-2-2m0-4v-2m0-4a2 2 0 0 1 2-2m4 0h2m4 0a2 2 0 0 1 2 2m0 4v2m0 4v4m2-2h-4m-4 0h-2"/>`),
		g.Group(children),
	)
}