package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M15 6H9a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1ZM9 4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3H9Z" clip-rule="evenodd"/><path d="M2 6a1 1 0 0 1 2 0v12a1 1 0 1 1-2 0V6Zm18 0a1 1 0 1 1 2 0v12a1 1 0 1 1-2 0V6Z"/></g>`),
		g.Group(children),
	)
}