package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneQuickreply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4v13.17L5.17 16H15v-6h5V4z" opacity=".3"/><path fill="currentColor" d="M5.17 16L4 17.17V4h16v6h2V4c0-1.1-.9-2-2-2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h9v-2H5.17z"/><path fill="currentColor" d="m19 23l3.5-7h-2.2l1.7-4h-5v6h2z"/>`),
		g.Group(children),
	)
}