package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAmigo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 12a2 2 0 1 0 4 0a2 2 0 1 0-4 0"/><path d="m9.591 3.635l-7.13 14.082c-1.712 3.38 1.759 5.45 3.69 3.573l1.86-1.81c3.142-3.054 4.959-2.99 8.039.11l1.329 1.337c2.372 2.387 5.865.078 4.176-3.225L14.36 3.635c-1.114-2.18-3.666-2.18-4.77 0z"/></g>`),
		g.Group(children),
	)
}