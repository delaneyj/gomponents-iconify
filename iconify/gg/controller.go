package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Controller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="m14.828 6.343l1.415-1.414L12 .686L7.757 4.93l1.415 1.414L12 3.515l2.828 2.828Zm-9.899 9.9l1.414-1.415L3.515 12l2.828-2.828L4.93 7.757L.686 12l4.243 4.243Zm2.828 2.828L12 23.314l4.243-4.243l-1.415-1.414L12 20.485l-2.828-2.828l-1.415 1.414Zm9.9-9.899L20.485 12l-2.828 2.828l1.414 1.415L23.314 12L19.07 7.757l-1.414 1.415Z"/><path fill-rule="evenodd" d="M12 8a4 4 0 1 1 0 8a4 4 0 0 1 0-8Zm0 2a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}