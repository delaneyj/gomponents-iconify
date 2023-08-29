package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M4.293 6.293A1 1 0 0 0 5 8h2v6a1 1 0 1 0 2 0V8h2a1 1 0 0 0 .707-1.707l-3-3a1 1 0 0 0-1.414 0l-3 3zm8 11.414A1 1 0 0 1 13 16h2v-6a1 1 0 1 1 2 0v6h2a1 1 0 0 1 .707 1.707l-3 3a1 1 0 0 1-1.414 0l-3-3z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}