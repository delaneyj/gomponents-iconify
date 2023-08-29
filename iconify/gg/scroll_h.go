package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M7.182 9.172L5.768 7.757L1.525 12l4.243 4.243l1.414-1.415L4.353 12l2.829-2.828Zm9.636 5.656l1.414 1.415L22.475 12l-4.243-4.243l-1.414 1.415L19.646 12l-2.828 2.828Z"/><path fill-rule="evenodd" d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm-2 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}