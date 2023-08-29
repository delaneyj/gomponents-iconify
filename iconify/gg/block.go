package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Block(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8.465 14.121a1 1 0 1 0 1.414 1.415l5.657-5.657a1 1 0 1 0-1.415-1.415l-5.656 5.657Z"/><path fill-rule="evenodd" d="M6.343 17.657A8 8 0 1 0 17.657 6.343A8 8 0 0 0 6.343 17.657Zm9.9-1.414a6 6 0 1 1-8.486-8.485a6 6 0 0 1 8.486 8.485Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}