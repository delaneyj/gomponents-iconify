package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M11 5a1 1 0 0 0-1 1h4a1 1 0 0 0-1-1h-2zm0-2a3 3 0 0 0-3 3H4a1 1 0 0 0 0 2h1v10a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8h1a1 1 0 1 0 0-2h-4a3 3 0 0 0-3-3h-2zm0 8a1 1 0 1 0-2 0v5a1 1 0 1 0 2 0v-5zm4 0a1 1 0 1 0-2 0v5a1 1 0 1 0 2 0v-5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}