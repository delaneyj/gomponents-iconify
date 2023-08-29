package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 8l4 8m-4 0l4-8m4 0v8h2a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2z"/>`),
		g.Group(children),
	)
}