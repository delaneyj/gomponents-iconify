package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdjustmentsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M8 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM5 6a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm11 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0zm-5 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0z"/><path d="M3 6a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H4a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1zm-6 6a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1zm14 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1zM3 18a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1zm6 0a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1z"/></g>`),
		g.Group(children),
	)
}