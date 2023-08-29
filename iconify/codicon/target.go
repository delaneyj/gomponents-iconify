package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Target(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path d="M8 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z"/><path d="M12 8a4 4 0 1 1-8 0a4 4 0 0 1 8 0Zm-4 3a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M15 8A7 7 0 1 1 1 8a7 7 0 0 1 14 0Zm-7 6A6 6 0 1 0 8 2a6 6 0 0 0 0 12Z"/></g>`),
		g.Group(children),
	)
}