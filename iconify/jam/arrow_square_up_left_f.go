package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSquareUpLeftF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.172 6.172a.997.997 0 0 0-1 1v6a1 1 0 0 0 2 0V9.586l3.95 3.95a1 1 0 0 0 1.414-1.415l-3.95-3.95h3.586a1 1 0 0 0 0-2h-6zM4 0h12a4 4 0 0 1 4 4v12a4 4 0 0 1-4 4H4a4 4 0 0 1-4-4V4a4 4 0 0 1 4-4z"/>`),
		g.Group(children),
	)
}