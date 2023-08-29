package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Designmodo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7 4.054a5 5 0 1 0 0 10a5 5 0 0 0 0-10Zm-2 5a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z" clip-rule="evenodd"/><path d="M22 10.554h-5v-3h5v3Zm-10.45 6.392a3.988 3.988 0 0 0 2.829-1.172l2.121 2.121a6.978 6.978 0 0 1-4.95 2.05a6.978 6.978 0 0 1-4.95-2.05l2.122-2.12a3.987 3.987 0 0 0 2.828 1.17Z"/></g>`),
		g.Group(children),
	)
}