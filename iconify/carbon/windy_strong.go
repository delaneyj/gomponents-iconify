package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindyStrong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13 30a5.006 5.006 0 0 1-5-5h2a3 3 0 1 0 3-3H4v-2h9a5 5 0 0 1 0 10Z"/><path fill="currentColor" d="M25 25a5.006 5.006 0 0 1-5-5h2a3 3 0 1 0 3-3H2v-2h23a5 5 0 0 1 0 10zm-4-13H6v-2h15a3 3 0 1 0-3-3h-2a5 5 0 1 1 5 5z"/>`),
		g.Group(children),
	)
}