package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M10 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM9 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1-7a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm4-7a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 10a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm2-13a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm1 2a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-1 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path fill-rule="evenodd" d="M8 1a3 3 0 0 0-3 3v16a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V4a3 3 0 0 0-3-3H8Zm8 2H8a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}