package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableColAfter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0a2 2 0 0 1 2 2v1a1 1 0 0 1-2 0V2H6v10h6v-1a1 1 0 0 1 2 0v1a2 2 0 0 1-2 2H2a2 2 0 0 1-2-2V2a2 2 0 0 1 2-2h10ZM4 10H2v2h2v-2Zm6-6a1 1 0 0 1 1 1v1h1a1 1 0 0 1 0 2h-1v1a1 1 0 0 1-2 0V8H8a1 1 0 1 1 0-2h1V5a1 1 0 0 1 1-1ZM4 6H2v2h2V6Zm0-4H2v2h2V2Z"/>`),
		g.Group(children),
	)
}