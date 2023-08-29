package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutlineMarkEmailRead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h8v-2H4V8l8 5l8-5v5h2V6c0-1.1-.9-2-2-2zm-8 7L4 6h16l-8 5zm5.34 11l-3.54-3.54l1.41-1.41l2.12 2.12l4.24-4.24L23 16.34L17.34 22z"/>`),
		g.Group(children),
	)
}