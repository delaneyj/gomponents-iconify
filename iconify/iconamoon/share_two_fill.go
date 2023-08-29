package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 14a1 1 0 1 0 2 0V6h2a1 1 0 0 0 .707-1.707l-3-3a1 1 0 0 0-1.414 0l-3 3A1 1 0 0 0 9 6h2v8ZM4 8a1 1 0 0 0-1 1v9a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V9a1 1 0 0 0-1-1h-2a1 1 0 1 0 0 2h1v8a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-8h1a1 1 0 1 0 0-2H4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}