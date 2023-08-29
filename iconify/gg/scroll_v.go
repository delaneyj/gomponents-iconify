package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScrollV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m9.172 16.818l-1.415 1.414L12 22.475l4.243-4.243l-1.415-1.414L12 19.647l-2.828-2.829Zm5.656-9.636l1.415-1.414L12 1.525L7.757 5.768l1.415 1.414L12 4.354l2.828 2.828Z"/><path fill-rule="evenodd" d="M12 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6Zm0 2a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}