package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BuildingInsightsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 2H16a2.002 2.002 0 0 0-2 2v10H4a2.002 2.002 0 0 0-2 2v14h28V4a2.002 2.002 0 0 0-2-2ZM9 28v-7h4v7Zm19 0H15v-8a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v8H4V16h12V4h12Z"/><path fill="currentColor" d="M18 8h2v2h-2zm6 0h2v2h-2zm-6 6h2v2h-2zm6 0h2v2h-2zm-6 6h2v2h-2zm6 0h2v2h-2zM9 12H7a5.006 5.006 0 0 1 5-5v2a3.003 3.003 0 0 0-3 3zm-5 0H2A10.011 10.011 0 0 1 12 2v2a8.01 8.01 0 0 0-8 8z"/>`),
		g.Group(children),
	)
}