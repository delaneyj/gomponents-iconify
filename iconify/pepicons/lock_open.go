package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 7H7a4 4 0 0 0-4 4v4a4 4 0 0 0 4 4h6a4 4 0 0 0 4-4v-4a4 4 0 0 0-4-4Zm-8 4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-4Z" clip-rule="evenodd"/><path d="M8 8a1 1 0 1 1-2 0V4a4 4 0 1 1 8 0v.5a1 1 0 1 1-2 0V4a2 2 0 1 0-4 0v4Z"/><path fill-rule="evenodd" d="M8 13a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm2.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}