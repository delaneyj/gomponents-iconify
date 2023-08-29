package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.859 6H5.025A2.025 2.025 0 0 0 3 8.025v2.834c0 .537.213 1.052.593 1.432l6.116 6.116a2.025 2.025 0 0 0 2.864 0l2.834-2.834a2.025 2.025 0 0 0 0-2.864L9.29 6.593A2.025 2.025 0 0 0 7.859 6z"/><path d="m17.573 18.407l2.834-2.834a2.025 2.025 0 0 0 0-2.864L13.29 5.593M6 9h-.01"/></g>`),
		g.Group(children),
	)
}