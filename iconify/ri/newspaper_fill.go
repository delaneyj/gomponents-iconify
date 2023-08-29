package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewspaperFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a3 3 0 0 1-3-3V3a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v7h4v9a3 3 0 0 1-3 3Zm-1-10v7a1 1 0 1 0 2 0v-7h-2ZM5 6v6h6V6H5Zm0 7v2h10v-2H5Zm0 3v2h10v-2H5Zm2-8h2v2H7V8Z"/>`),
		g.Group(children),
	)
}