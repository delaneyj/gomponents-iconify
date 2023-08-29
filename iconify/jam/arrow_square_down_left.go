package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4 2a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H4zm0-2h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/><path d="m8.172 10.414l3.95-3.95a1 1 0 1 1 1.414 1.415l-3.95 3.95h3.586a1 1 0 0 1 0 2h-6a.997.997 0 0 1-1-1v-6a1 1 0 1 1 2 0v3.585z"/></g>`),
		g.Group(children),
	)
}