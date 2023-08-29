package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertColumnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 3a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h6Zm-1 2h-4v14h4V5ZM6 7a5 5 0 1 1 0 10A5 5 0 0 1 6 7Zm1 2H5v1.999L3 11v2l2-.001V15h2v-2.001L9 13v-2l-2-.001V9Z"/>`),
		g.Group(children),
	)
}