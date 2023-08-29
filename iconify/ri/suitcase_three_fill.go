package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SuitcaseThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 1a1 1 0 0 1 1 1v5h1V6h2v1h1a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-1v1h-2v-1H7v1H5v-1H4a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h1V6h2v1h1V2a1 1 0 0 1 1-1h6Zm-6 9H7v8h2v-8Zm4 0h-2v8h2v-8Zm4 0h-2v8h2v-8Zm-3-7h-4v4h4V3Z"/>`),
		g.Group(children),
	)
}