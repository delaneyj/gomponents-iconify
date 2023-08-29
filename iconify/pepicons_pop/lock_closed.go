package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockClosed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 13a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm2.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/><path d="M13 7H7a4 4 0 0 0-4 4v4a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-4a4 4 0 0 0-4-4Zm-8 4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-4Z"/><path d="M14 4a4 4 0 0 0-8 0v5h8V4Zm-2 0v3H8V4a2 2 0 1 1 4 0Z"/></g>`),
		g.Group(children),
	)
}