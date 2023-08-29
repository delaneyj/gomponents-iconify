package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11 5a1 1 0 0 0-1.707-.707l-7 7a1 1 0 0 0 0 1.414l7 7A1 1 0 0 0 11 19v-3.025c1.691-.011 3.83.133 5.633.583c1.088.27 1.973.633 2.565 1.076c.567.424.802.864.802 1.366a1 1 0 1 0 2 0c0-1.925-.598-4.66-2.42-6.937c-1.719-2.15-4.462-3.805-8.58-4.036V5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}