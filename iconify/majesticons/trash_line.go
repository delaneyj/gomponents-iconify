package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrashLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 7a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2v10a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V8a1 1 0 0 1-1-1zm3 1v10a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1V8H7z"/><path d="M11 5a1 1 0 0 0-1 1v1a1 1 0 0 1-2 0V6a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v1a1 1 0 1 1-2 0V6a1 1 0 0 0-1-1h-2zm-1 5a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1zm4 0a1 1 0 0 1 1 1v5a1 1 0 1 1-2 0v-5a1 1 0 0 1 1-1z"/></g>`),
		g.Group(children),
	)
}