package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M2 7a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2zm0 5a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2zm2 3a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4H4zm5-8a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4h-2a2 2 0 0 1-2-2zm2 3a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4h-2zm-2 7a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4h-2a2 2 0 0 1-2-2zm9-12a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4h-2zm-2 7a2 2 0 0 1 2-2h2a2 2 0 1 1 0 4h-2a2 2 0 0 1-2-2zm2 3a2 2 0 1 0 0 4h2a2 2 0 1 0 0-4h-2z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}