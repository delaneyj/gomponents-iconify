package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectionLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M9 9h6m-3 6V9"/><path d="M6 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0ZM22 4a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm0 16a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path stroke-linecap="round" d="M6 20h12m0-16H6m14 14V6M4 6v12" opacity=".5"/></g>`),
		g.Group(children),
	)
}