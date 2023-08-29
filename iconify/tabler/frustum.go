package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frustum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m18.402 5.508l2.538 10.158a1.99 1.99 0 0 1-1.064 2.278L12.84 21.31a1.945 1.945 0 0 1-1.682 0l-7.035-3.365a1.99 1.99 0 0 1-1.064-2.278L5.598 5.508a1.98 1.98 0 0 1 1.11-1.328l4.496-2.01a1.95 1.95 0 0 1 1.59 0l4.496 2.01c.554.246.963.736 1.112 1.328z"/><path d="m18 4.82l-5.198 2.324a1.963 1.963 0 0 1-1.602 0L6 4.819m6 2.501V21.5"/></g>`),
		g.Group(children),
	)
}