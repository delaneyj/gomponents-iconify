package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h4.586L2 28.586L3.414 30L10 23.414V28h2v-8H4zm26-10l-6-6l-1.414 1.414L27.172 10l-4.586 4.586L24 16l6-6zm-16.08 7.484l4.15-15.483l1.932.517l-4.15 15.484zM4 10l6-6l1.414 1.414L6.828 10l4.586 4.586L10 16l-6-6z"/>`),
		g.Group(children),
	)
}