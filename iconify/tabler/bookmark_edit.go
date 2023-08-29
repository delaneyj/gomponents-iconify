package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkEdit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.35 17.39L7 20V6a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v6m1.42 3.61a2.1 2.1 0 0 1 2.97 2.97L18 22h-3v-3l3.42-3.39z"/>`),
		g.Group(children),
	)
}